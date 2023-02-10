// Code generated by hertz generator.

package core

import (
	"context"

	core "github.com/ClubWeGo/douyin/biz/model/core"
	"github.com/ClubWeGo/douyin/kitex_server"
	"github.com/ClubWeGo/videomicro/kitex_gen/videomicro"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// FeedMethod .
// @router /douyin/feed/ [GET]
func FeedMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.FeedReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	msgsucceed := "获取视频流成功"
	msgFailed := "获取视频流失败"

	resp := new(core.FeedResp)

	// 本项目没有采用时间作为查询依据，使用offset作为依据，
	// 存在问题：如果后续有新视频上传，会导致给用户推荐重复视频流 明天修改，微服务改成使用时间查询的方式

	vclient := kitex_server.Videoclient
	r, err := vclient.GetVideosFeedMethod(context.Background(), &videomicro.GetVideosFeedReq{Offset: 0, Limit: 30})
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = &msgFailed
		c.JSON(consts.StatusOK, resp)
		return
	}

	resp.VideoList = make([]*core.Video, 0)
	for _, video := range r.VideoList {
		author, _ := kitex_server.GetUser(video.AuthorId)
		// 暂时不做处理，错误返回空对象即可
		resp.VideoList = append(resp.VideoList, &core.Video{
			ID:            video.Id,
			Author:        author,
			PlayURL:       video.PlayUrl,
			CoverURL:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false, // 需要增加喜欢配置
			Title:         video.Title,
		})
	}

	resp.StatusMsg = &msgsucceed

	c.JSON(consts.StatusOK, resp)
}
