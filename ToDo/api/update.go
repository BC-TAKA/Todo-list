package api

import (
	"log"

	"github.com/raveger/Todo-list/ToDo/common"
)

//update.go---UPDATE用の構造体
type UpdateData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Todo string `json:"todo"`
}

//UPDATE処理を行う関数
func Update(todo UpdateData) {
	db := common.DbConn()
	_, err := db.Exec(
		`UPDATE todolist SET name=?, TODO=? WHERE id=?`,
		todo.Name, todo.Todo, todo.ID,
	)
	if err != nil {
		log.Fatal(err)
	}
}
