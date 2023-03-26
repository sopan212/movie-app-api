package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/movie-app-api/auth"
	"github.com/movie-app-api/config"
	"github.com/movie-app-api/models"
)

func GetUser(c *gin.Context) {
	var users []models.User

	config.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"message": "succes",
		"data":    users,
	})
}
func RegisterUser(c *gin.Context) {
	var user models.User
	var profile models.Profile
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}
	err = c.BindJSON(&profile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}

	err = user.Hashpassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed hash password",
			"err":     err.Error(),
		})
		c.Abort()
		return
	}
	user.Profile = profile
	insertUser := config.DB.Create(&user)
	if insertUser.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed insert data",
			"error":   insertUser.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        user.ID,
		"user name": user.Username,
		"email":     user.Email,
	})
}

func GenerateToken(c *gin.Context) {
	request := models.TokenRequest{}
	user := models.User{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
			"err":     err.Error(),
		})
		c.Abort()
		return
	}

	//check email

	checkEmail := config.DB.Where("email = ?", request.Email).First(&user)
	fmt.Println("email", request)
	if checkEmail.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "email not found",
			"error":   checkEmail.Error.Error(),
		})
		c.Abort()
		return
	}

	//check password

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauhtorized",
			"error":   credentialError.Error(),
		})
		c.Abort()
		return

	}
	//generate token

	tokenString, err := auth.GenerateJTW(user.Email, user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "internal server error",
			"err":     err.Error(),
		})
		c.Abort()
		return
	}
	//response
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}
