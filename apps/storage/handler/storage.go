package handler

import (
	"context"
	"io"
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
	db.Raw("SELECT tablename FROM pg_tables WHERE schemaname = 'public'").Scan(&tables)

	for _, table := range tables {
		fmt.Println(table)
	}
}

func NewPostgresRInstance(host string) *gorm.DB {
	dsn := fmt.Sprintf("host=postgres user=jian password=123456 dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Shanghai")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Errorf("NewPostgresRInstance err %s", err.Error())
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	//sqlDB.SetMaxIdleConns(viper.GetInt("POSTGRES_MAX_IDLE"))
	//sqlDB.SetMaxOpenConns(viper.GetInt("POSTGRES_OPEN_IDLE"))
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}