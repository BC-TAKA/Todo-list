package api

import (
	"encoding/json"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/raveger/Todo-list/ToDo/common"
)

type getData struct {
	Name string `json:"name"`
	Todo string `json:"TODO"`
}

//DBに登録を行う関数
func Ins() {
	readBody, err := ioutil.ReadAll(r.body)
	//JSONをデコード
	var datas []getData
	if err := json.Unmarshal(readBody, &datas); err != nil {
		log.Fatal(err)
	}

	db := common.DbConn()
	_, err := db.Exec(
		`INSERT INTO todolist(name,TODO) VALUES(?,?)`,
		datas.Name, datas.TODO,
	)
	if err != nil {
		log.Fatal(err)
	}
}
