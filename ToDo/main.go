package main

import (
	"./api"
)

//type TodoList struct {
//	Name string `json:"name"`
//	Todo string `json:"todo"`
//}

func main() {
	insert := api.NewIns()

}

//func メソッド名(w http.ResponseWriter, r *http.Request){
//	実行させたい処理内容
//	実行させたい処理内容
//}

//全件表示の関数
//func Index(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	selDB, err := db.Query("SELECT * FROM todolist")
//	if err != nil {
//		panic(err.Error())
//	}
//	td := todolist{}
//	res := []todolist{}
//	for selDB.Next() {
//		var name, TODO string
//		err = selDB.Scan(&name, &TODO)
//		if err != nil {
//			panic(err.Error())
//		}
//		td.name = name
//		td.TODO = TODO
//		res = append(res, td)
//	}

//}

//htmlで入力されたデータをconversion.js経由でDBに登録する関数
//func jsondecoding(w http.ResponseWriter, r *http.Request) {
//	bytes, err := ioutil.ReadFile("conversion.js")
//	if err != nil {
//		log.Fatal(err)
//	}
//	var todolist []TodoList
//	if err := json.Unmarshal(bytes, &todolist); err != nil {
//		log.Fatal(err)
//	}
//}

//html側で入力されたTodoをformからPOST形式で取得
//func jsonEncoding(w http.ResponseWriter, r *http.Request){
//	v := r.FormValue("ValueTodo")
//	todoData, err := json.Marshal(v)
//この後どうDBに繋げていくのか
//}

//func handleRequests() {
//http.HandleFunc("/", jsondecoding)
//	http.HandleFunc("/", jsonEncoding)
//	log.Fatal(http.ListenAndServe(":8081", nil))
//}

//func main() {
//	handleRequests()
//}
