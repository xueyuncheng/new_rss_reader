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
	feed := router.Group("/feeds")
	{
		feed.POST("", wrap(AddFeed))
		feed.GET("", wrap(GetFeed))
		feed.DELETE("/:id", wrap(DeleteFeed))
	}

	news := router.Group("/news")
	{
		news.GET("", wrap(GetNews))
	}
}

func wrap(f func(c *gin.Context) interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		d := f(c)
		c.JSON(200, d)
	}
}
