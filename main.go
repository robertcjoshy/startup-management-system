package main

import (
	"E/golang_gin_postgresql/start_up/src/models"
	//"E/golang_gin_postgresql/start_up/src/middlewares"
	"E/golang_gin_postgresql/start_up/src/routes"
	//"E/golang_gin_postgresql/start_up/src/utils"
)

func main() {
	//utils.LoadEnv()
	models.OpenDatabaseConnection()
	models.AutoMigrateModels()
	router := routes.SetupRoutes()
	router.Run("localhost:8080")

}
