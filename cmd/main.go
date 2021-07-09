package main

import (
	"log"

	"github.com/Hudayberdyyev/todo-app-Maksim-"
	"github.com/Hudayberdyyev/todo-app-Maksim-/pkg/handler"
	"github.com/Hudayberdyyev/todo-app-Maksim-/pkg/repository"
	"github.com/Hudayberdyyev/todo-app-Maksim-/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	// handlers := new(handler.Handler)

	if err := initConfig(); err != nil {
		log.Fatalf("error config: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
