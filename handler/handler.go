package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}


func CreateOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "POST Opening",
	})
}

func UpdateOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "PUT Opening",
	})
}

func DestroyOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "DELETE Opening",
	})
}

func ListOpeningsHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "GET Openings",
	})
}

