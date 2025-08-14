package routes

import (
	"go-url-shrtr/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func shorten(context *gin.Context) {
	var url model.Url

	if err := context.ShouldBindJSON(&url); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data: " + err.Error()})
		return
	}

	if err := url.Save(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not save url: " + err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "url short code created successfully ✅"})
}
