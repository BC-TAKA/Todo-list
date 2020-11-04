package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	)

type TodoList struct {
	Id int 'json:"id"'
	Name string 'json:"name"'
	Todo string 'json:"todo"'
}

//func メソッド名(w http.ResponseWriter, r *http.Request){
//	実行させたい処理内容
//	実行させたい処理内容
//}

//htmlで入力されたデータをconversion.js経由でDBに登録する関数
func jsondecoding(w http.ResponseWriter, r *http.Request){
	bytes,err := ioutil.ReadFile("conversion.js")
	if err != nil {
		log.Fatal(err)
	}

	var todolist []TodoList
	if err := json.Unmarshal(bytes, &todolist); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", "user:password@tcp(host:port/todo")
	if err := db.QueryRow("SELECT * FROM todo_list"); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, err)
}

//html側で入力されたTodoをformからPOST形式で取得
//func jsonEncoding(w http.ResponseWriter, r *http.Request){
//	v := r.FormValue("ValueTodo")
//	todoData, err := json.Marshal(v)
//この後どうDBに繋げていくのか
//}

func handleRequests(){
	http.HandleFunc("/", jsondecoding)
//	http.HandleFunc("/", jsonEncoding)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
