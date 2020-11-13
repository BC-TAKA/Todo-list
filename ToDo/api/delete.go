package api

import (
	"log"

	"github.com/raveger/Todo-list/ToDo/common"
)

func Del(id int) {
	db := common.DbConn()
	if _, err := db.Exec("DELETE FOROM todolist WHERE id=?", id); err != nil {
		log.Fatal(err)
	}
}
