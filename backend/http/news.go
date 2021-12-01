package http

import (
	"myNewFeed/internal/ecode"
	"myNewFeed/model"
	"myNewFeed/service"

	"github.com/gin-gonic/gin"
)

func ListNews(ctx *gin.Context) interface{} {
	req := &model.ListNewsReq{}
	if err := ctx.ShouldBind(req); err != nil {
		return ecode.ErrInternal.WithData(err.Error())
	}

	news, err := service.ListNews(ctx, req)
	if err != nil {
		return ecode.ErrInternal.WithData(err.Error())
	}

	newsResp := make([]*model.ListNewsResp, 0, len(news))
	for _, v := range news {
		tmp := &model.ListNewsResp{
			ID:          v.ID,
			Title:       v.Title,
			Link:        v.Link,
			PublishTime: v.PublishTime.Format("2006-01-02 15:04:05"),
			FeedID:      v.FeedID,
			FeedName:    v.FeedName,
		}

		newsResp = append(newsResp, tmp)
	}

	return ecode.ErrOK.WithData(newsResp)
}

func StatNews(ctx *gin.Context) interface{} {
	chart, err := service.StatNews(ctx)
	if err != nil {
		return ecode.ErrInternal.WithData(err.Error())
	}

	return ecode.ErrOK.WithData(chart)
}
