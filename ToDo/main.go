package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/raveger/Todo-list/ToDo/api"
)

// func listInsert(w http.ResponseWriter, r *http.Request) {
// 	api.Ins()
// }

func listEncode(w http.ResponseWriter, r *http.Request) {
	todos := api.List()
	ecd := json.NewEncoder(w)
	ecd.Encode(&todos)
}

func main() {
	//http.Handleでサーバー接続
	http.Handle("/", http.FileServer(http.Dir(".")))
	//	http.HandleFunc("/insert", listInsert)
	http.HandleFunc("/todos", listEncode)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
