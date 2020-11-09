package api

import (
	"log"

	"github.com/raveger/Todo-list/ToDo/common"
)

type Listup struct {
	ID   int
	Name string
	Todo string
}

func List() []Listup {
	var id int
	var name string
	var TODO string
	baseList := []Listup{}

	db := common.DbConn()
	rows, err := db.Query("SELECT * FROM todolist")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&id, &name, &TODO)
		if err != nil {
			log.Fatal(err)
		}
		base := Listup{
			ID:   id,
			Name: name,
			Todo: TODO,
		}
		baseList = append(baseList, base)
	}
	return baseList
}
