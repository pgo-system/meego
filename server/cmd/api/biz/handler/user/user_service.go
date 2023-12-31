// Code generated by hertz generator.

package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	base "p_meego/server/cmd/api/biz/model/base"
	user "p_meego/server/cmd/api/biz/model/user"
	"p_meego/server/cmd/api/config"
	"p_meego/server/shared/errno"
	kuser "p_meego/server/shared/kitex_gen/user"
)

// Login .
// @router /login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.LoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	hlog.Debugf("Login request: %v", req)

	// 根据用户名查询用户信息
	infoRequest := kuser.UserInfoRequest{}
	infoRequest.SetType("user_name")
	infoRequest.SetKey(req.GetUserName())

	info, err := config.GlobalUserProviderClient.UserInfo(ctx, &infoRequest)

	if err != nil {
		c.JSON(consts.StatusOK, user.LoginResponse{
			Code:    errno.BusinessError,
			Message: err.Error(),
		})
		return
	}

	if info.GetCode() != errno.Success {
		c.JSON(consts.StatusOK, user.LoginResponse{
			Code:    info.GetCode(),
			Message: info.GetMessage(),
		})
		return
	}

	userInfo := info.GetData()
	uid := userInfo.GetUid()

	// 根据用户信息中的uid用于登陆
	loginRequest := kuser.LoginRequest{}
	loginRequest.SetUid(uid)
	loginRequest.SetPassword(req.GetPassword())

	loginResp, err := config.GlobalUserProviderClient.Login(ctx, &loginRequest)
	if err != nil {
		c.JSON(consts.StatusOK, user.LoginResponse{
			Code:    errno.BusinessError,
			Message: err.Error(),
		})
		return
	}

	if loginResp.GetCode() != errno.Success {
		c.JSON(consts.StatusOK, user.LoginResponse{
			Code:    loginResp.GetCode(),
			Message: loginResp.GetMessage(),
		})
		return
	}

	resp := new(errno.Response)
	resp.Success(loginResp.GetData().GetUid())

	c.JSON(consts.StatusOK, resp)
}

// Register .
// @router /register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.RegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, user.LoginResponse{
			Code:    errno.BusinessError,
			Message: err.Error(),
		})
		return
	}
	registerRequest := kuser.RegisterRequest{}
	registerRequest.SetEmail(req.GetEmail())
	registerRequest.SetPassword(req.GetPassword())
	registerRequest.SetUserName(req.GetUserName())
	registerRequest.SetAvatarUrl(req.GetAvatarURL())
	registerRequest.SetPhoneNumber(req.GetPhoneNumber())
	registerRequest.SetPersonalSignature(req.GetPersonalSignature())
	r, err := config.GlobalUserProviderClient.Register(ctx, &registerRequest)
	if err != nil {
		c.JSON(consts.StatusOK, user.LoginResponse{
			Code:    errno.BusinessError,
			Message: err.Error(),
		})
		return
	}
	var uid int32
	if r.GetData() != nil {
		uid = r.GetData().GetUid()
	}
	resp := user.LoginResponse{
		Code:    r.GetCode(),
		Message: r.GetMessage(),
		Data:    uid,
	}

	c.JSON(consts.StatusOK, resp)
}

// Logout .
// @router /logout [POST]
func Logout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.LogoutRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.BaseResponse)

	c.JSON(consts.StatusOK, resp)
}
