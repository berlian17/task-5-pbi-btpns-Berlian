package controllers

import (
	"net/http"
	"task-5-pbi-btpns-Berlian/models"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email" valid:"email,required"`
	Password string `json:"password" valid:"required,length(6|255)"`
}

func Login(c *gin.Context) {
	var input LoginInput

	// Mengikat input JSON ke struct input.
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi input menggunakan govalidator.
	if _, err := govalidator.ValidateStruct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	user.Email = input.Email
	user.Password = input.Password

	token, err := models.LoginCheck(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "email or password is incorect"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login success",
		"token": token,
	})
}

type RegisterInput struct {
	Username string `json:"username" valid:"required"`
	Email    string `json:"email" valid:"email,required"`
	Password string `json:"password" valid:"required,length(6|255)"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	// Mengikat input JSON ke struct input.
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi input menggunakan govalidator.
	if _, err := govalidator.ValidateStruct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password

	_, err := user.SaveUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration success"})
}
