package controllers

import (
	"go-gin-api/config"
	"net/http"
	"time"

	models "go-gin-api/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := models.User{
		ID:        "fsdlfsdkfsd",
		Title:     "Manager",
		Body:      "The body of the user",
		Completed: "Not completed",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	conn := config.ConnectFirestore()
	models.AddToCollection(user, &conn)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Users",
		"data":    user,
	})
}

func GetAllUsers(c *gin.Context) {
	user := models.User{}
	conn := config.ConnectFirestore()
	users := models.FindAll(user, &conn)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Users",
		"data":    users,
	})
}
