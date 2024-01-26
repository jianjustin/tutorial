package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/lib/pq"
)

func hello(w http.ResponseWriter, r *http.Request) {
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

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":3000", nil)
}
