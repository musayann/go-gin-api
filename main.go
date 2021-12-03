package main

import (
	"log"

	"github.com/gin-gonic/gin"

	routes "go-gin-api/routes"
)

func main() {

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}
