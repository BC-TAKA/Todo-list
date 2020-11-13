package api

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/raveger/Todo-list/ToDo/common"
)

type GetData struct {
	Name string `json:"name"`
	Todo string `json:"todo"`
}

//DBに登録を行う関数
func Ins(todo GetData) {
	log.Println(todo)

	db := common.DbConn()
	_, err := db.Exec(
		`INSERT INTO todolist(name,TODO) VALUES(?,?)`,
		todo.Name, todo.Todo,
	)
	if err != nil {
		log.Fatal(err)
	}
}
