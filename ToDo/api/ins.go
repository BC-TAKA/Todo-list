package api

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//DBに登録を行う関数
func ins() {
	db := conn.dbConn()
	insDB, err = db.Exec(
		`INSERT INTO todolist(name,TODO) VALUES(?,?)`,
		"test11", "testTodo11",
	)
	if err != nil {
		log.Fatal(err)
	}
	return
}
