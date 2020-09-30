package controllers

import (
	"api-practice2/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var todos = []models.Todo{
	{ID: "1", Name: "Take garbage out"},
	{ID: "2", Name: "Wash windows"},
	{ID: "3", Name: "Search for a job"},
}


func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	fmt.Println("hello")
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("There was a problem decoding"))
		return
	}

	todos = append(todos, todo)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(todo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("there was a problem encoding"))
		return
	}
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("there was a problem encoding"))
	}

}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Problem converting ID to integer"))
		return
	}

	todo := todos[id - 1]

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Problem encoding value"))
		return
	}

}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Problem converting ID to integer"))
		return
	}

	var updatedTodo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Problem decoding the data"))
	}

	todos[id - 1] = updatedTodo

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("There was a problem encoding the data"))
	}
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("There was a problem converting ID to integer"))
	}

	todos = append(todos[:id - 1], todos[id:]...)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("There was a problem encoding the data"))
	}
}
