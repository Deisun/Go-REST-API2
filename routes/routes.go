package routes

import (
	"api-practice2/controllers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {

	router.HandleFunc("/todos", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/todos", controllers.GetAllTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", controllers.GetTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", controllers.DeleteTodo).Methods("DELETE")
}
