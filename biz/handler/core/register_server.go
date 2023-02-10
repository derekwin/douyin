// Code generated by hertz generator.

package core

import (
	"context"

	core "github.com/ClubWeGo/douyin/biz/model/core"
	"github.com/ClubWeGo/douyin/kitex_server"
	"github.com/ClubWeGo/douyin/tools"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// RegisterMethod .
// @router /douyin/user/register/ [POST]
func RegisterMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	msgsucceed := "注册成功"
	msgFailed := "注册失败"

	resp := new(core.RegisterResp)

	userid, err := kitex_server.RegisterUser(req.Username, *req.Password)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = &msgFailed
		c.JSON(consts.StatusOK, resp)
		return
	}

	resp.StatusMsg = &msgsucceed
	resp.Token = tools.GetToken(userid)
	resp.UserID = userid

	c.JSON(consts.StatusOK, resp)
}
