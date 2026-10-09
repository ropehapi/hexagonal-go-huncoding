package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ropehapi/hexagonal-hun-coding/adapter/input/controller"
	"github.com/ropehapi/hexagonal-hun-coding/adapter/output/news_http"
	"github.com/ropehapi/hexagonal-hun-coding/application/service"
)

func InitRoutes(r *gin.Engine) {
	newsClient := news_http.NewNewsClient()
	newsService := service.NewNewsService(newsClient)
	newsController := controller.NewNewsController(newsService)

	r.GET("/news", newsController.GetNews)
}