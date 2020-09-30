package main

import (
	"api-practice2/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", r))

}
