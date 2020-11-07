package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/raveger/Todo-list/ToDo/api"
)

func listEncode(w http.ResponseWriter, r *http.Request) {
	todos := api.List()
	ecd := json.NewEncoder(w)
	ecd.Encode(&todos)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/todos", listEncode)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
