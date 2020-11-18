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

//INSERT処理を行う関数
func CreateTODO(todo GetData) {
	db := common.DbConn()
	_, err := db.Exec(
		`INSERT INTO todolist(name,TODO) VALUES(?,?)`,
		todo.Name, todo.Todo,
	)
	if err != nil {
		log.Fatal(err)
	}
}
