package model

//main.go---INSERT用の構造体
type GetData struct {
	Name string `json:"name"`
	Todo string `json:"todo"`
}

//main.go---UPDATE用の構造体
type UpdateData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Todo string `json:"todo"`
}
