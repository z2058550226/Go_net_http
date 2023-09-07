package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/get", handleGet)
	http.HandleFunc("/postJson", handlePostJson)
	http.HandleFunc("/postForm", handlePostForm)
	http.HandleFunc("/responseJson", handleResponseJson)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	fmt.Fprintf(w, "GET: id = %s \n", id)
}

func handlePostJson(w http.ResponseWriter, r *http.Request) {
	// 根据请求body创建一个json解析器实例
	decoder := json.NewDecoder(r.Body)
	// 用于存放参数key=value数据
	var param map[string]string
	// 解析参数 存入map
	decoder.Decode(&param)
	fmt.Fprintf(w, "POST json: username=%s, password=%s\n", param["username"], param["password"])
}

func handlePostForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username:=r.Form.Get("username")
	password:=r.Form.Get("password")
	fmt.Fprintf(w, "Post form: username=%s, password=%s\n", username, password)
}

// response json data
func handleResponseJson(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Code int
		Msg  string
		Data interface{}
	}
	res := Response{
		200,
		"success",
		"admin",
	}
	json.NewEncoder(w).Encode(res)
}
