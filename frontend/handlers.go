package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jianjustin/frontend/proto/add"
	pb "github.com/jianjustin/frontend/proto/mul"
	"github.com/jianjustin/frontend/proto/sub"
	_ "github.com/lib/pq"
	"io"
	"net/http"
)

type requestBody struct {
	A int64 `json:"a"`
	B int64 `json:"b"`
}

func (fe *frontendServer) AddHandler(w http.ResponseWriter, r *http.Request) {
	var body requestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, _ := fe.addService.Add(r.Context(), &add.AddRequest{A: body.A, B: body.B})

	fmt.Fprint(w, result.GetResult())
}

func (fe *frontendServer) MulHandler(w http.ResponseWriter, r *http.Request) {
	var body requestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, _ := fe.mulService.Mul(r.Context(), &pb.MulRequest{A: body.A, B: body.B})

	fmt.Fprint(w, result.GetResult())
}

func (fe *frontendServer) SubHandler(w http.ResponseWriter, r *http.Request) {
	var body requestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, _ := fe.subService.Sub(r.Context(), &sub.SubRequest{A: body.A, B: body.B})

	fmt.Fprint(w, result.GetResult())
}

func (fe *frontendServer) HelloHandler(w http.ResponseWriter, r *http.Request) {
	// PostgreSQL数据库连接信息
	const (
		host     = "postgres" // 使用服务名称作为主机名
		port     = 5432
		user     = "jian"
		password = "123456"
		dbname   = "testdb"
	)

	// 构建连接`字符串
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// 打开数据库连接
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	defer db.Close()

	// 检查数据库连接是否成功
	err = db.Ping()
	if err != nil {
		//log.Fatal(err)
		io.WriteString(w, err.Error())
		return
	}

	fmt.Println("Successfully connected!")
	io.WriteString(w, "Successfully connected!")
}
