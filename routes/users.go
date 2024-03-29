package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"org.gfoo/api-rest/models"
	"org.gfoo/api-rest/utils"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"error": "Could not parse request data (not fit the model)"})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "Could not save user in DB"})
		return
	}

	context.JSON(http.StatusCreated,
		gin.H{"message": "User created successfully!"},
	)
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"error": "Could not parse request data (not fit the model)"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized,
			gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "Could not generate token"})
		return
	}

	context.JSON(http.StatusOK,
		gin.H{"message": "User logged in successfully!", "token": token})
}
