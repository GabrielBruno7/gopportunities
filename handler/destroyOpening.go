package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DestroyOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "DELETE Opening",
	})
}
