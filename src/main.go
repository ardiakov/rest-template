package main

import (
	"./utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	utils.SuccessResponse([]string{"Main"}, w)
}

func withPathParams(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if userId, ok := params["userId"]; ok {
		utils.SuccessResponse([]string{userId}, w)
	}
}

func withQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	log.Println(query)

	utils.SuccessResponse([]string{""}, w)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/path-query/{userId}", withPathParams)
	router.HandleFunc("/query", withQuery)

	log.Fatal(http.ListenAndServe(":8080", router))
}
