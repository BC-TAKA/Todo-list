package model

//全件表示用の構造体
//todolistテーブルでIDをint型としているため、TodoData構造体では一覧表示が出来ないので残しています。
type Listup struct {
	ID   int
	Name string
	Todo string
}

//全件表示機能以外で使用する構造体
type TodoData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Todo string `json:"todo"`
}
