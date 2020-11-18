package api

import (
	"log"

	"github.com/raveger/Todo-list/ToDo/common"
	"github.com/raveger/Todo-list/ToDo/model"
)

//UPDATE処理を行う関数
func Update(todo model.UpdateData) {
	db := common.DbConn()
	_, err := db.Exec(
		`UPDATE todolist SET name=?, TODO=? WHERE id=?`,
		todo.Name, todo.Todo, todo.ID,
	)
	if err != nil {
		log.Fatal(err)
	}
}
