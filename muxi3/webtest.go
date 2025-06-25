package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func startServer() {
	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("title")
		if title == "" {
			title = "空白"
		}
		fmt.Fprintf(w, "您正在查询图书《 %s》", title)
	}) //Get请求
	http.HandleFunc("/comment", func(w http.ResponseWriter, r *http.Request) {
		type C struct {
			User    string `json:"user"`
			Comment string `json:"comment"`
		}
		var p C
		//从请求体中解析json到结构体
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			fmt.Printf("JSON解析失败: %v\n", err)
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		// 设置响应头为 JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// 返回 JSON 响应
		type Response struct {
			Message string `json:"message"`
			User    string `json:"user"`
			Comment string `json:"comment"`
		}
		json.NewEncoder(w).Encode(Response{
			Message: "评论提交成功",
			User:    p.User,
			Comment: p.Comment,
		}) //Post接口
	})
	fmt.Println("服务启动于http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
func handleGet() {
	resp, err := http.Get("http://127.0.0.1:8080/book?title=三体")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("响应内容(Get)", string(body)) //处理Get请求
}
func handlePost() {
	data := map[string]interface{}{
		"user":    "小明",
		"comment": "这本书真棒",
	}
	jsonData, _ := json.Marshal(data)
	//发送POST请求
	resp, err := http.Post("http://localhost:8080/comment",
		"application/json", bytes.NewReader(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("响应内容(POST)", string(body)) //处理post请求
}
func runClient() {
	time.Sleep(100 * time.Millisecond)
	handleGet()
	handlePost()
}
func main() {
	go startServer()
	runClient()

	select {}
}
