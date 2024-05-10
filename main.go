/* To-Do API */

package main

import (
	"log"
	"net/http"

	"github.com/UjjwalMahar/gotodo/database"
	handlers "github.com/UjjwalMahar/gotodo/internal/api"
)
func main(){

	log.Print("starting application ...")

	database.Init()

	//Handler func 
	http.HandleFunc("/new-todo", handlers.NewTodo) //post

	http.HandleFunc("/get-todo", handlers.GetTodo) //get

	http.HandleFunc("/update-todo", handlers.UpdateTodo) //update

	http.HandleFunc("/delete-todo", handlers.DeleteTodo) //delete

	http.ListenAndServe(":4000",nil)
}
