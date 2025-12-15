package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Coldwws/kolesa/internal/handler"
	"github.com/Coldwws/kolesa/internal/repository"
	"github.com/Coldwws/kolesa/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatalf("Ошибка подключения к бд: ", err)
	}

	storage := repository.NewCarPostgres(db)
	carService := service.NewCarService(storage)
	handler := handler.NewHandler(carService)

	router := handler.InitRoutes()


	server := http.Server{
		Addr:    ":9090",
		Handler: router,
	}
	fmt.Println("Server started at port: 9090")
	server.ListenAndServe()
}
