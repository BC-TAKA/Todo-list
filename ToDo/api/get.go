package api

import (
	"log"

	"github.com/raveger/Todo-list/ToDo/common"
)

//最後にこれだけを実施
type Listup struct {
	ID   int
	Name string
	Todo string
}

func GetTODOs() []Listup {
	var (
		id   int
		name string
		TODO string
	)

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
