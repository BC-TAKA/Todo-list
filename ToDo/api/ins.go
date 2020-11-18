package api

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/raveger/Todo-list/ToDo/common"
	"github.com/raveger/Todo-list/ToDo/model"
)

//INSERT処理を行う関数
func CreateTODO(todo model.GetData) {
	db := common.DbConn()
	_, err := db.Exec(
		`INSERT INTO todolist(name,TODO) VALUES(?,?)`,
		todo.Name, todo.Todo,
	)
	if err != nil {
		log.Fatal(err)
	}
}
