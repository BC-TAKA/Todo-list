package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/raveger/Todo-list/ToDo/api"
)

//INSERT用の構造体
type GetData struct {
	Name string `json:"name"`
	Todo string `json:"todo"`
}

//UPDATE用の構造体
type UpdateData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Todo string `json:"todo"`
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	var todo UpdateData
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	log.Println(todo)
}

//delete.go
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	var id int
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.Error(w, "ID not found", 400)
		return
	}
	api.Del(id)
}

//ins.go
func insertTodo(w http.ResponseWriter, r *http.Request) {
	var todo api.GetData
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	api.Ins(todo)
}

//get.go
func listEncode(w http.ResponseWriter, r *http.Request) {
	todos := api.List()
	ecd := json.NewEncoder(w)
	ecd.Encode(&todos)
}

func main() {
	//http.Handleでサーバー接続
	http.Handle("/", http.FileServer(http.Dir(".")))
	// http.HandleFunc("/update", updateTodo)
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		//機能はrequestの内容で分岐
		switch r.Method {
		case http.MethodGet:
			listEncode(w, r)
		case http.MethodPost:
			insertTodo(w, r)
		case http.MethodDelete:
			deleteTodo(w, r)
		case http.MethodPut:
			updateTodo(w, r)
		}
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
