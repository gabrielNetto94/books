package main

import (
	"books/internal/adapters/rest/routes"
	"log"
)

func main() {

	if err := routes.InitRoutes().Run(":3000"); err != nil {
		log.Fatal("Error init server", err.Error())
	}

}
