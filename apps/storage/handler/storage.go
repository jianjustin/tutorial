package handler

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/types/known/structpb"
	"gorm.io/gorm"
	"io"
	"storage/middleware"
	"storage/model"
	"time"

	"go-micro.dev/v4/logger"

	pb "storage/proto"
)

type Storage struct {
	DB *gorm.DB
}

func NewStorageHandler(db *gorm.DB) *Storage {
	return &Storage{DB: db}
}

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
	db := middleware.NewPostgresInstance()
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

func (e *Storage) ConnectCache(ctx context.Context, req *pb.ConnectCacheRequest, rsp *pb.ConnectCacheResponse) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	list, err := rdb.Keys(ctx, "*").Result()
	if err != nil {
		return err
	}

	keys := []*structpb.Struct{}

	for _, key := range list {
		keyStruct, err := structpb.NewStruct(map[string]interface{}{
			"key": key,
		})
		if err != nil {
			return err
		}
		keys = append(keys, keyStruct)
	}
	rsp.Datas = keys
	return nil
}

func (e *Storage) Create(ctx context.Context, req *pb.CreateRequest, rsp *pb.CreateResponse) error {
	tenantId, err := middleware.GetTenantId(ctx)
	if err != nil {
		return err
	}

	dataMap := req.Record.GetFields()
	logger.Infof("Received Storage.Create request: %v", req)
	systemParam := model.SystemParam{
		Id:    cast.ToString(dataMap["id"].GetStringValue()),
		Key:   cast.ToString(dataMap["key"].GetStringValue()),
		Value: cast.ToString(dataMap["value"].GetStringValue()),
		G:     tenantId,
	}
	err = e.DB.Create(&systemParam).Error
	if err != nil {
		return err
	}

	return nil
}
