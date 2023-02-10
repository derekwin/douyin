// Code generated by hertz generator.

package core

import (
	"context"

	core "github.com/ClubWeGo/douyin/biz/model/core"
	"github.com/ClubWeGo/douyin/kitex_server"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// PublishListMethod .
// @router /douyin/publish/list/ [GET]
func PublishListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.PublishListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	msgsucceed := "获取用户视频列表成功"
	msgFailed := "获取用户视频列表失败"

	resp := new(core.PublishListResp)

	videolist, err := kitex_server.GetVideosByAuthorId(req.UserID)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = &msgFailed
		c.JSON(consts.StatusOK, resp)
		return
	}
	resp.VideoList = videolist
	resp.StatusMsg = &msgsucceed
	c.JSON(consts.StatusOK, resp)
}
