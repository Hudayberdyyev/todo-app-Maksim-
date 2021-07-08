package main

import (
	"log"

	"github.com/Hudayberdyyev/todo-app-Maksim-"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error http server: %s", err.Error())
	}

}
