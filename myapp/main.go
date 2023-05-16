package main

import (
	"myapp/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.InitializeRoutes()
}
