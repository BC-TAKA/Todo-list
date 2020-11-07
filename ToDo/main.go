package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/raveger/Todo-list/ToDo/api"
)

func listEncode(w http.ResponseWriter, r *http.Request) {
	todos := api.List()
	ecd := json.NewEncoder(w)
	ecd.Encode(&todos)
}

func handleRequests() {
	http.HandleFunc("/", listEncode)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
	//	api.Ins()
	// todos := api.List()
	// ecd := json.NewEncoder(os.Stdout)
	// if err := ecd.Encode(todos); err != nil {
	// 	fmt.Println(err)
	// }
}
