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

func Search(id int) {
	var searchId int
	var searchName string
	var searchTodo string
	searchList := []SearchElements{}

	db := common.DbConn()
	rows, err := db.Query("SELECT * FROM todolist WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&id, &searchName, &searchTodo)
		if err != nil {
			log.Fatal(err)
		}
		base := SearchElements{
			ID:   id,
			Name: searchName,
			Todo: searchTodo,
		}
		SearchResult = append(SearchResult, base)
	}
	return SearchResult
}
