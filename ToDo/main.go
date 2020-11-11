package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/raveger/Todo-list/ToDo/api"
)

type getData struct {
	Name string `json:"name"`
	Todo string `json:"todo"`
}

func test(w http.ResponseWriter, r *http.Request) {
	var todo getData
	if err := json.NewDecoder(r.body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func listEncode(w http.ResponseWriter, r *http.Request) {
	//ここで分岐させるように修正
	todos := api.List()
	ecd := json.NewEncoder(w)
	ecd.Encode(&todos)
}

func main() {
	//http.Handleでサーバー接続
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/todos", test)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
