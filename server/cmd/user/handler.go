package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/crypto/bcrypt"
	"p_meego/server/cmd/user/pkg/model"
	"p_meego/server/cmd/user/pkg/mysql"
	"p_meego/server/shared/errno"
	"p_meego/server/shared/kitex_gen/base"
	"p_meego/server/shared/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	UserContext *mysql.Query
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.UserResponse, err error) {
	resp = user.NewUserResponse()
	queryCtx := s.UserContext.User
	do := s.UserContext.WithContext(ctx).User
	do = do.Where(queryCtx.UID.In(req.GetUid()))
	find, err := do.Find()

	if err != nil {
		klog.Error("Login uid db error ", req.GetUid())
		dbError(resp)
		return resp, nil
	}

	if len(find) == 0 {
		klog.Error("Login uid not found ", req.GetUid())
		resp.SetMessage("用户不存在")
		noExistError(resp)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(find[0].Password), []byte(req.GetPassword()))
	if err != nil {
		klog.Error("Login uid pass word error ", req.GetUid())
		resp.SetMessage("密码错误")
		paramsError(resp)
		return resp, nil
	}

	resp = success(find[0])
	return
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.UserResponse, err error) {
	resp = user.NewUserResponse()

	queryCtx := s.UserContext.User
	do := s.UserContext.WithContext(ctx).User
	do = do.Where(queryCtx.UserName.In(req.GetUserName()))
	find, err := do.Find()
	if err != nil {
		dbError(resp)
		return resp, nil
	}
	if len(find) != 0 {
		resp.SetMessage("用户名已存在")
		noExistError(resp)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		klog.Error(ctx, "encrypt password")
	}

	userModel := model.User{
		UserName:          req.GetUserName(),
		Email:             req.GetEmail(),
		AvatarURL:         req.GetAvatarUrl(),
		PersonalSignature: req.GetPersonalSignature(),
		PhoneNumber:       req.GetPhoneNumber(),
		Password:          string(hashedPassword),
	}

	// 用户新建
	err = do.Create(&userModel)

	// 返回详情
	userInfoReq := user.UserInfoRequest{
		Type: "user_name",
		Key:  req.GetUserName(),
	}
	return s.UserInfo(ctx, &userInfoReq)
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserResponse, err error) {
	resp = user.NewUserResponse()
	key := req.GetKey()
	keyType := req.GetType()
	if keyType != "user_name" {
		resp.SetMessage("不支持的类型")
		paramsError(resp)
		klog.Error("UserInfo error params ", req.GetKey())
		return
	}

	queryCtx := s.UserContext.User
	do := s.UserContext.WithContext(ctx).User

	do = do.Where(queryCtx.UserName.In(key))
	find, err := do.Find()
	if err != nil {
		dbError(resp)
		klog.Error("UserInfo db error params ", req.GetKey())
		return resp, nil
	}
	if len(find) == 0 {
		resp.SetMessage("用户信息不存在")
		noExistError(resp)
		klog.Error("UserInfo no existed error params ", req.GetKey())
		return
	}

	resp = success(find[0])

	return
}

func success(userModel *model.User) (resp *user.UserResponse) {
	userInfo := base.UserInfo{}
	userInfo.SetUserName(userModel.UserName)
	userInfo.SetUid(userModel.UID)
	userInfo.SetEmail(userModel.Email)
	userInfo.SetAvatarUrl(userModel.AvatarURL)
	userInfo.SetPersonalSignature(userModel.PersonalSignature)
	userInfo.SetPhoneNumber(userModel.PhoneNumber)
	userInfo.SetPoints(userModel.Points)
	resp = user.NewUserResponse()
	resp.SetCode(errno.Success)
	resp.SetMessage(errno.SuccessMessage)
	resp.SetData(&userInfo)
	return
}

func paramsError(resp *user.UserResponse) {
	resp.SetCode(errno.ParamsError)
	if resp.GetMessage() != "" {
		return
	}
	resp.SetMessage(errno.ParamsErrorMessage)
}

func dbError(resp *user.UserResponse) {
	resp.SetCode(errno.DatabaseError)
	if resp.GetMessage() != "" {
		return
	}
	resp.SetMessage(errno.DatabaseErrorMessage)
}

func noExistError(resp *user.UserResponse) {
	resp.SetCode(errno.NoExistError)
	if resp.GetMessage() != "" {
		return
	}
	resp.SetMessage(errno.NoExistErrorMessage)
}
