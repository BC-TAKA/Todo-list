package api

import (
	"log"

	"github.com/raveger/Todo-list/ToDo/common"
)

//DELETE処理を行う関数
func DeleteTODO(id int) {
	db := common.DbConn()
	if _, err := db.Exec("DELETE FROM todolist WHERE id=?", id); err != nil {
		log.Fatal(err)
	}
}
