package main

import (
	"log"

	"github.com/Hudayberdyyev/todo-app-Maksim-"
	"github.com/Hudayberdyyev/todo-app-Maksim-/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error http server: %s", err.Error())
	}

}
