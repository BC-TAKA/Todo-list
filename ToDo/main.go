package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/raveger/Todo-list/ToDo/api"
)

func insertTodo(w http.ResponseWriter, r *http.Request) {
	insert := api.Ins()
	dcd := json.NewDecoder(r.body)
	dcd.Decode(&insert)
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
			//POST形式でメソッドがきたらこの分岐
		case http.MethodPost:
			insertTodo(w, r)
		}
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
