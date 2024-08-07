package router

import (
	"github.com/GabrielBruno7/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	{
		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.DELETE("/opening", handler.DestroyOpeningHandler)

		v1.GET("/openings", handler.ListOpeningsHandler)
	}

}
