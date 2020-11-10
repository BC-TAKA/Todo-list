package api

import (
	"encoding/json"
	"fmt"
	"log"
	"syscall/js"

	_ "github.com/go-sql-driver/mysql"
	"github.com/raveger/Todo-list/ToDo/common"
)

type CheckStruct struct {
	Name string `json:"name"`
	TODO string `json:"todo"`
}

func unmarshalJSON(src js.Value, dst interface{}) {
	str := js.Global().Get("JSON").Call("stringify", src).String()
	err := json.Unmarshal([]byte(str), &dst)
	if err != nil {
		fmt.Println(err)
	}
}

func CheckJSON(this js.Value, args []js.Value) interface{} {
	var checkStruct CheckStruct
	unmarshalJSON(args[0], &checkStruct)
	fmt.Println(checkStruct)
	return nil
}

func registerCallbacks() {
	js.Global().Set("CheckJson", js.FuncOf(CheckJSON))
}

//DBに登録を行う関数
func Ins() {
	db := common.DbConn()
	_, err := db.Exec(
		`INSERT INTO todolist(name,TODO) VALUES(?,?)`,
		"test22", "testTodo22",
	)
	if err != nil {
		log.Fatal(err)
	}
	return
}
