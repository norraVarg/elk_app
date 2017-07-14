package main

import (
	"net/http"

	"./api"
	"github.com/gorilla/mux"
)

func main() {
	// Initial state
	model := new(api.Model)
	model.Players = make(map[string]api.Player)
	model.Ch = make(chan int, 1)
	model.Ch <- 0

	router := mux.NewRouter()
	router.HandleFunc("/player/{id:[0-9]+}", model.Update)
	router.HandleFunc("/statistic", model.GetStatistic)

	http.ListenAndServe(":8080", router)
}
