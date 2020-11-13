package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/raveger/Todo-list/ToDo/api"
)

type GetData struct {
	Name string `json:"name"`
	Todo string `json:"todo"`
}

type SelectData struct {
	id int
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "no id", 400)
		return
	}
	fmt.Println(id)
	//	api.Del(id)
}

func insertTodo(w http.ResponseWriter, r *http.Request) {
	var todo api.GetData
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	api.Ins(todo)
}

func listEncode(w http.ResponseWriter, r *http.Request) {
	todos := api.List()
	ecd := json.NewEncoder(w)
	ecd.Encode(&todos)
}

func main() {
	//http.Handleでサーバー接続
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			listEncode(w, r)
		case http.MethodPost:
			insertTodo(w, r)
		case http.MethodDelete:
			deleteTodo(w, r)
		}
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
