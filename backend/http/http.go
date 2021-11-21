package http

import (
	"fmt"
	"myNewFeed/internal/log"
	"myNewFeed/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitHttp(config *model.Http) {
	router := gin.Default()
	router.Use(cors.Default())
	InitRouter(router)
	if err := router.Run(fmt.Sprintf(":%v", config.Port)); err != nil {
		log.Sugar.Fatalw("http服务启动失败", "err", err)
	}
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
