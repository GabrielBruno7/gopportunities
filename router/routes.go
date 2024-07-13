package router

import (
	"github.com/GabrielBruno7/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	{
		v1.GET("/opening", func(context *gin.Context) {
			handler.ShowOpeningHandler(context)
		})

		v1.POST("/opening", func(context *gin.Context) {
			handler.CreateOpeningHandler(context)
		})

		v1.PUT("/opening", func(context *gin.Context) {
			handler.UpdateOpeningHandler(context)
		})

		v1.DELETE("/opening", func(context *gin.Context) {
			handler.DestroyOpeningHandler(context)
		})

		v1.GET("/openings", func(context *gin.Context) {
			handler.ListOpeningsHandler(context)
		})

	}

}
