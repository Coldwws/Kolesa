package main

import (
	"fmt"
	"net/http"
	"github.com/Coldwws/kolesa/internal/handler"
)

func main() {
	handler := handler.NewHandler()

	router := handler.InitRoutes()
	server := http.Server{
		Addr: ":9090",
		Handler: router,
	}
	fmt.Println("Server started at port: 9090")
	server.ListenAndServe()
}