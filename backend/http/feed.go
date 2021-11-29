package http

import (
	"myNewFeed/internal/ecode"
	"myNewFeed/model"
	"myNewFeed/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListFeed(ctx *gin.Context) interface{} {
	feeds, err := service.ListFeed(ctx)
	if err != nil {
		return ecode.ErrInternal.WithData(err.Error())
	}

	feedsResp := make([]*model.ListFeedResp, 0, len(feeds))
	for _, feed := range feeds {
		tmp := &model.ListFeedResp{
			ID:       feed.ID,
			Name:     feed.Name,
			ErrorMsg: feed.ErrorMsg,
		}
		feedsResp = append(feedsResp, tmp)
	}

	return ecode.ErrOK.WithData(feedsResp)
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
