/*
This file contains Todo struct definition. It represents the data structure of a todo item.
*/

package model

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
}

