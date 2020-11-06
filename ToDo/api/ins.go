package api

import (
	"./common"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Ins struct{}

//DBに登録を行う関数
func ins() {
	con := common.NewConn()
	db := con.DbConn()

	insDB, err := db.Exec(
		`INSERT INTO todolist(name,TODO) VALUES(?,?)`,
		"test11-2", "testTodo11-2",
	)
	if err != nil {
		log.Fatal(err)
	}
	return
}
