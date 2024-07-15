package api

import (
	_ "api/api/docs"
	"api/api/handler"
	"api/api/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @version 1.0
// @title User
// @description API Gateway
// @host localhost:8080
// BasePath: /
func Router(hand *handler.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	stories := router.Group("/api/v1/stories")
	stories.Use(middleware.Check)
	{
		stories.POST("/:story_id", hand.CreateStory)
		stories.PUT("/:story_id")
		stories.DELETE("/:story_id")
		stories.GET("")
		stories.GET("/:story_id")
		stories.POST("/:story_id/comments")
		stories.GET("/:story_id/comments")
		stories.POST("/:story_id/like")
	}
	itineraries := router.Group("/api/v1/itineraries")
	itineraries.Use(middleware.Check)
	{
		itineraries.POST("")
		itineraries.PUT("/:itinerary_id")
		itineraries.DELETE("/:itinerary_id")
		itineraries.GET("")
		itineraries.GET("/:itinerary_id")
		itineraries.POST("/:itinerary_id/comments")

	}
	return router
}
