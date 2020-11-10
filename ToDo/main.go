package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/raveger/Todo-list/ToDo/api"
)

//新規登録のエンドポイントも/todosにするため
//listEncodeでの処理内容をHTTPメソッド参照で分岐させる必要がある。
//兼平さんのサンプル参照
func listEncode(w http.ResponseWriter, r *http.Request) {
	todos := api.List()
	ecd := json.NewEncoder(w)
	ecd.Encode(&todos)
}

func main() {
	//http.Handleでサーバー接続
	http.Handle("/", http.FileServer(http.Dir(".")))
	//	http.HandleFunc("/insert", listInsert)
	http.HandleFunc("/todos", listEncode)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
