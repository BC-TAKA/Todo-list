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
	ID   string `json:"id"`
	Name string `json:"name"`
	Todo string `json:"todo"`
}

//update.go
func updateTodo(w http.ResponseWriter, r *http.Request) {
	var todo api.UpdateData
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		log.Println(err)
		return
	}
	api.Update(todo)
}

//delete.go
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	//登録しているIDが1からなので、1より小さい数値をエラーとして検出する
	if err != nil || id < 1 {
		log.Println(err)
		return
	}
	api.DeleteTODO(id)
}

//ins.go
func createTODO(w http.ResponseWriter, r *http.Request) {
	var todo api.GetData
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		log.Println(err)
		return
	}
	api.CreateTODO(todo)
}

//get.go
func getTODOs(w http.ResponseWriter, r *http.Request) {
	todos := api.GetTODOs()
	ecd := json.NewEncoder(w)
	ecd.Encode(&todos)
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
