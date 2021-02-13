package main

import (
	"log"

	"github.com/Valeriy-Totubalin/myface-go"
	"github.com/Valeriy-Totubalin/myface-go/internal/delivery/http"
)

func main() {
	handlers := new(http.Handler)

	srv := new(myface.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http sever: %s", err.Error())
	}
}
