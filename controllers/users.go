package controllers

import (
	"net/http"
	"time"

	models "go-gin-api/models"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users := []models.User{
		models.User{
			ID:        "fsdlfsdkfsd",
			Title:     "Manager",
			Body:      "The body of the user",
			Completed: "Not completed",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Users",
		"data":    users,
	})
}
