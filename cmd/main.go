package main

import (
	"log"
	todo_list "todo-list"
	"todo-list/internal/handler"
)

//set endpoints
//connect database
//create repository
//create service
//create handler
//init dependency injection

func main() {
	handlers := new(handler.Handler)
	srv := new(todo_list.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error with running server: %s", err.Error())
	}
}
