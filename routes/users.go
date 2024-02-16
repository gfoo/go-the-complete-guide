package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"org.gfoo/api-rest/models"
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
