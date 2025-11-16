package routes

import (
	"net/http"

	"github.com/afrique-ai-be/models"
	"github.com/gin-gonic/gin"
)

func saveNewsLetterEmail(context *gin.Context) {
	var users models.UserEmails
	err := context.ShouldBindJSON(&users)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data!"})
		return
	}
	err = users.Save()

	if err != nil {
		if err.Error() == "email has already subscribed" {
			context.JSON(http.StatusConflict, gin.H{"message": "email has already subscribed"})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user!"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Successfully subscribed to newletter!"})
}
