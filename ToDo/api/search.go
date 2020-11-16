package api

import (
	"log"

	"github.com/raveger/Todo-list/ToDo/common"
)

type SearchElements struct {
	ID   int
	Name string
	Todo string
}

func Search(id int) []SearchElements {
	var searchId int
	var name string
	var TODO string
	SearchResult := []SearchElements{}

	db := common.DbConn()
	rows, err := db.Query("SELECT * FROM todolist WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&searchId, &name, &TODO)
		if err != nil {
			log.Fatal(err)
		}
		base := SearchElements{
			ID:   searchId,
			Name: name,
			Todo: TODO,
		}
		SearchResult = append(SearchResult, base)
	}
	log.Println(SearchResult)
	return SearchResult
}
