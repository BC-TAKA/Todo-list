package common

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Conn struct{}

type db struct{}

func GetConn() {

}

func CloseConn() {

}

//DBに接続するための関数
func dbConn() {
	var db *sql.DB
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/todo")
	if err != nil {
		panic(err.Error())
	}
	return db
}
