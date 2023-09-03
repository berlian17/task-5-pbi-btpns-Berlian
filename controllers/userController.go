package controllers

import (
	"net/http"
	"task-5-pbi-btpns-Berlian/models"
	"task-5-pbi-btpns-Berlian/utils/token"

	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.GetUserByID(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": user,
	})
}