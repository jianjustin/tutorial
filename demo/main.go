package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// 定义一个结构体来存储 POST 请求的数据
type User struct {
	A int64 `json:"a"`
	B int64 `json:"b"`
}

func main() {
	// 创建一个新的路由器
	r := mux.NewRouter()

	// 添加路由处理程序
	r.HandleFunc("/post", PostHandler).Methods("POST")

	// 启动 HTTP 服务器并指定路由器
	log.Fatal(http.ListenAndServe(":8080", r))
}

// 处理 POST 请求的处理程序
func PostHandler(w http.ResponseWriter, r *http.Request) {
	// 解析 JSON 请求体
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	// 打印解析后的结构体内容
	fmt.Fprintf(w, "Received POST request!\n")
	fmt.Fprintf(w, "Username: %d\n", user.A)
	fmt.Fprintf(w, "Password: %d\n", user.B)
}
