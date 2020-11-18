package api

import (
	"log"

	"github.com/raveger/Todo-list/ToDo/common"
)

func DeleteTODO(id int) {
	log.Println(id)
	db := common.DbConn()
	if _, err := db.Exec("DELETE FROM todolist WHERE id=?", id); err != nil {
		log.Fatal(err)
	}
}
