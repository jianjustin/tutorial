package embeds

import (
	"fmt"
	"io"
	"net/http"

	"go.guide/tutorial/embeds/resources"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// 使用 embed.FS.Open() 函数打开 index.html 文件
	file, err := resources.StaticFiles.Open("resources/index.html")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// 读取文件内容并写入 ResponseWriter
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(data)
}

func Server() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/index", indexHandler)

	http.ListenAndServe(":8090", nil)
}
