package router

import (
	"Hacktiv8/Golang/Assignment/Assignment_3/handler"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/weather")
	{
		userRouter.GET("/status", handler.GetStatusHandler)
	}
	return router
}
