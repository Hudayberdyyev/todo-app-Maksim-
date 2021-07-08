package main

import (
	"log"

	"github.com/Hudayberdyyev/todo-app-Maksim-"
	"github.com/Hudayberdyyev/todo-app-Maksim-/pkg/handler"
	"github.com/Hudayberdyyev/todo-app-Maksim-/pkg/repository"
	"github.com/Hudayberdyyev/todo-app-Maksim-/pkg/service"
)

func main() {
	// handlers := new(handler.Handler)
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error http server: %s", err.Error())
	}

}
