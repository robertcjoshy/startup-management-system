package routes

import (
	"E/golang_gin_postgresql/start_up/src/controllers"

	"github.com/gin-gonic/gin"
)

func startupsGroupRouter(baseRouter *gin.RouterGroup) {
	startups := baseRouter.Group("/startups")

	startups.GET("/all", controllers.GetAllStartups)
	startups.GET("/get/:id", controllers.GetStartupByID)
	startups.POST("/create", controllers.CreateStartup)
	startups.PATCH("/update", controllers.UpdateStartup)
	startups.DELETE("/delete/:id", controllers.DeleteStartup)
}

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	versionRouter := r.Group("")
	startupsGroupRouter(versionRouter)

	return r
}
