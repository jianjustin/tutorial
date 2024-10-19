package handler

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/structpb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"os"
	"time"

	"go-micro.dev/v4/logger"

	pb "storage/proto"
)

type Storage struct{}

func (e *Storage) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received Storage.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *Storage) ClientStream(ctx context.Context, stream pb.Storage_ClientStreamStream) error {
	var count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			logger.Infof("Got %v pings total", count)
			return stream.SendMsg(&pb.ClientStreamResponse{Count: count})
		}
		if err != nil {
			return err
		}
		logger.Infof("Got ping %v", req.Stroke)
		count++
	}
}

func (e *Storage) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.Storage_ServerStreamStream) error {
	logger.Infof("Received Storage.ServerStream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		logger.Infof("Sending %d", i)
		if err := stream.Send(&pb.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *Storage) BidiStream(ctx context.Context, stream pb.Storage_BidiStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		logger.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

func (e *Storage) Connect(ctx context.Context, req *pb.ConnectRequest, rsp *pb.ConnectResponse) error {
	db := NewPostgresRInstance()
	var tables []string
	err := db.Raw("SELECT tablename FROM pg_tables WHERE schemaname = 'public'").Scan(&tables).Error
	if err != nil {
		return err
	}

	type Table struct {
		Name string `json:"name"`
	}

	list := []*structpb.Struct{}
	for _, table := range tables {
		fmt.Println(table)
		tableStruct, err := structpb.NewStruct(map[string]interface{}{
			"name": table,
		})
		if err != nil {
			return err
		}
		list = append(list, tableStruct)
	}
	rsp.Datas = list
	return nil
}

func NewPostgresRInstance() *gorm.DB {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Errorf("NewPostgresRInstance err %s", err.Error())
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
