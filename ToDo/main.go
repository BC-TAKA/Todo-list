package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/raveger/Todo-list/ToDo/api"
	"github.com/raveger/Todo-list/ToDo/common"
	"github.com/raveger/Todo-list/ToDo/model"
)

//パッケージ変数DBを定義
var DB *sql.DB

//update.go
func updateTODO(w http.ResponseWriter, r *http.Request) {
	var todo model.TodoData
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	api.UpdateTODO(todo, DB)
}

//delete.go
func deleteTODO(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	api.DeleteTODO(id, DB)
}

//ins.go
func createTODO(w http.ResponseWriter, r *http.Request) {
	var todo model.TodoData
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	api.CreateTODO(todo, DB)
}

//get.go
func getTODOs(w http.ResponseWriter, r *http.Request) {
	todos := api.GetTODOs(DB)
	if err := json.NewEncoder(w).Encode(&todos); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func main() {
	//パッケージ変数DBにDB接続処理機能を代入
	DB = common.DbConn()
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
			deleteTODO(w, r)
		case http.MethodPut:
			updateTODO(w, r)
		}
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
