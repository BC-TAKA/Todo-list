package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	)

type todoList struct {
	Id int 'json:"id'
	Name string 'json:"name'
	Todo string 'json:"todo'
}

//func メソッド名(w http.ResponseWriter, r *http.Request){
//	実行させたい処理内容
//	実行させたい処理内容
//}


//html側で入力されたTodoをformからPOST形式で取得
func jsonEncoding(w http.ResponseWriter, r *http.Request){
	v := r.FormValue("ValueTodo")
	todoData, err := json.Marshal(v)
//この後どうDBに繋げていくのか
}

func handleRequests(){
	http.HandleFunc("/", jsonEncoding)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
