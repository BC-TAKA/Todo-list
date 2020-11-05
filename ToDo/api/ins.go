package api

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//DBに登録を行う関数
func ins() {
	db := api.dbConn()
	insDB, err = db.Exec(
		`INSERT INTO todolist(name,TODO) VALUES(?,?)`,
		"test11-2", "testTodo11-2",
	)
	if err != nil {
		log.Fatal(err)
	}
	return
}
