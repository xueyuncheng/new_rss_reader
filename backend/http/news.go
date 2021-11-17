package http

import (
	"myNewFeed/internal/ecode"
	"myNewFeed/service"

	"github.com/gin-gonic/gin"
)

func GetNews(ctx *gin.Context) interface{} {
	news, err := service.GetNews(ctx)
	if err != nil {
		return ecode.ErrInternal.WithData(err.Error())
	}

	return ecode.ErrOK.WithData(news)
}
