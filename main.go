package main

import "books/internal/adapters/rest/routes"

func main() {

	routes.InitRoutes().Run(":3000")

}
