package api

import (
	"../common"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//インスタンスするための空の構造体
type Ins struct{}

//DBに登録を行う関数
func ins() {
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
