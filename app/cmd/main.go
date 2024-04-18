package main

import (
	"log"

	"github.com/tibeahx/growers-dairy/app/api"
	"github.com/tibeahx/growers-dairy/app/storage"
	"github.com/tibeahx/growers-dairy/app/usecase"
)

func main() {
	db := storage.NewDB()
	serviceProvider := usecase.NewServiceProvider(db)
	handler := api.NewHandler(serviceProvider)
	server := api.NewServer(serviceProvider)
	if err := server.Run(":8080", handler.InitRoutes()); err != nil {
		log.Fatal("failed to start server", err.Error())
	}
}
