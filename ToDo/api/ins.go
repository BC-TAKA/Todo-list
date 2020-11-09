package api

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/raveger/Todo-list/ToDo/common"
)

//DBに登録を行う関数
func Ins() {
	db := common.DbConn()
	_, err := db.Exec(
		`INSERT INTO todolist(name,TODO) VALUES(?,?)`,
		"test22", "testTodo22",
	)
	if err != nil {
		log.Fatal(err)
	}
	return
}
