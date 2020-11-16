package api

import "log"

type UpdateData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Todo string `json:"todo"`
}

func Update(todo UpdateData) {
	log.Println(todo)

}
