package http

import (
	"myNewFeed/internal/ecode"
	"myNewFeed/model"
	"myNewFeed/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFeed(ctx *gin.Context) interface{} {
	feeds, err := service.GetFeed(ctx)
	if err != nil {
		return ecode.ErrInternal.WithData(err.Error())
	}

	return ecode.ErrOK.WithData(feeds)
}

func AddFeed(ctx *gin.Context) interface{} {
	req := &model.AddFeedReq{}
	if err := ctx.ShouldBind(req); err != nil {
		return ecode.ErrInvalidParams.WithData(err.Error())
	}

	if err := service.AddFeed(ctx, req); err != nil {
		return ecode.ErrInternal
	}

	return ecode.ErrOK
}

func DeleteFeed(ctx *gin.Context) interface{} {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ecode.ErrInvalidParams.WithData(err.Error())
	}

	if err := service.DeleteFeed(ctx, id); err != nil {
		return ecode.ErrInternal
	}

	return ecode.ErrOK
}
