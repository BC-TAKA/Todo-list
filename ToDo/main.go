package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/raveger/Todo-list/ToDo/api"
	"github.com/raveger/Todo-list/ToDo/model"
)

//update.go
func updateTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.UpdateData
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		log.Println(err)
		return
	}
	api.Update(todo)
}

//delete.go
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	//登録しているIDが1からなので、1より小さい数値もエラーとして検出する
	if err != nil || id < 1 {
		log.Println(err)
		return
	}
	api.DeleteTODO(id)
}

//ins.go
func createTODO(w http.ResponseWriter, r *http.Request) {
	var todo model.GetData
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		log.Println(err)
		return
	}
	api.CreateTODO(todo)
}

//get.go
func getTODOs(w http.ResponseWriter, r *http.Request) {
	todos := api.GetTODOs()
	if err := json.NewEncoder(w).Encode(&todos); err != nil {
		log.Println(err)
		return
	}
	// ecd := json.NewEncoder(w)
	// ecd.Encode(&todos)
}

func main() {
	//http.Handleでサーバー接続
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		//機能はrequestの内容で分岐
		switch r.Method {
		case http.MethodGet:
			getTODOs(w, r)
		case http.MethodPost:
			createTODO(w, r)
		case http.MethodDelete:
			deleteTodo(w, r)
		case http.MethodPut:
			updateTodo(w, r)
		}
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
