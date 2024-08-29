package main

import (
	"books/internal/adapters/rest/routes"
	"log"
)

func main() {

	// https://github.dev/sergicanet9/go-hexagonal-api

	if err := routes.InitRoutes().Run(":3000"); err != nil {
		log.Fatal("Error init server", err.Error())
	}

}
