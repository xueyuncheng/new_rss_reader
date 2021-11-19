package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitHttp() {
	router := gin.Default()
	router.Use(cors.Default())
	InitRouter(router)
	router.Run(":10001")
}

func InitRouter(router *gin.Engine) {
	api := router.Group("/api")

	feed := api.Group("/feeds")
	{
		feed.POST("", wrap(AddFeed))
		feed.GET("", wrap(ListFeed))
		feed.DELETE("/:id", wrap(DeleteFeed))
	}

	news := api.Group("/news")
	{
		news.GET("", wrap(ListNews))
	}
}

func wrap(f func(c *gin.Context) interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		d := f(c)
		c.JSON(200, d)
	}
}
