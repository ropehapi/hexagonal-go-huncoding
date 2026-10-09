package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ropehapi/hexagonal-hun-coding/adapter/input/routes"
	"github.com/ropehapi/hexagonal-hun-coding/configuration/logger"
)

func main(){
	logger.Info("About to start application")
	router := gin.Default()

	routes.InitRoutes(router)

	if err := router.Run(":8080"); err != nil {
		logger.Error("Error trying to init application on port 8080", err)
		return
	}
}