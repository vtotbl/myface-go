package main

import (
	"log"

	"github.com/Valeriy-Totubalin/myface-go"
	// "github.com/Valeriy-Totubalin/myface-go/pkg/handler"
	"github.com/Valeriy-Totubalin/myface-go/iternal/delivery/http/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(myface.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http sever: %s", err.Error())
	}
}
