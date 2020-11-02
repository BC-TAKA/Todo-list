//GOをここに記述

package main

import (
	"database/sql"
	"fmt"
	_ "gitHub.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"os"
)

type User struct {
	Id int
	Name string
	ToDo string
}

var idx *template.Template

func init() {
	idx = template.Must(template.ParseFiles("index.html"))
}

func main() {
	db, err := sql.Open("mysql", "root:password@/mydb")
	if err != nil {
		panic(err.Error())
	}

	defer db.close()

	rows, err != db.Query("SELECT * FROM todo_list")
	if err != nil {
		panic(err.Error())
	}
	
	for rows.Next() {
		var id int
		var name string
		var todo string
		if err := rows.Scan(&id, &name, &todo); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, todo)
	}
}
