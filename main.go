package main

import (
	"books/internal/adapters/rest/routes"
)

func main() {

	// https://github.dev/sergicanet9/go-hexagonal-api

	routes.InitRoutes().Run(":3000")

}
