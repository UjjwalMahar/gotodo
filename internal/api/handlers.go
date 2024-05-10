/*
This file contains HTTP handlers for API endpoints.
Each handler function will deal with a specific HTTP request,
like getting all to dos, creating a todo, updating a todo, etc.
*/

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/UjjwalMahar/gotodo/database"
	"github.com/UjjwalMahar/gotodo/internal/model"
)

//creating a newtodo handler

func NewTodo(w http.ResponseWriter, r *http.Request){
	var todo model.Todo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil{
		http.Error(w, "Failed to decode request body\n",http.StatusBadRequest)
		return
	}

		// Save the todo into MongoDB
	if err := database.SaveTodoToDB(todo); err != nil {
		http.Error(w, "Failed to save todo to database\n", http.StatusInternalServerError)
			return
	}
	
	fmt.Printf("Saved Successfully %v\n",http.StatusOK)

	w.WriteHeader(http.StatusCreated)


}

//get todo 
func GetTodo(w http.ResponseWriter, r *http.Request){
	
}
//delete todo 
func DeleteTodo(w http.ResponseWriter, r *http.Request){
	
}
//update todo
func UpdateTodo(w http.ResponseWriter, r *http.Request){
	
}
