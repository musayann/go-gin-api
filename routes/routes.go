package routes

import (
	"net/http"

	controllers "go-gin-api/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/users", controllers.GetAllUsers)
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
