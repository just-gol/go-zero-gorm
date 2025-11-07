// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"

	"zerogorm/internal/svc"
	"zerogorm/internal/types"
	"zerogorm/model/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddUserRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	user := gorm.User{
		Username: req.Username,
		Age:      req.Age,
		Email:    req.Email,
	}
	err = l.svcCtx.DB.Create(&user).Error
	if err != nil {
		return &types.CommonResponse{
			Code:    500,
			Message: "添加用户失败",
			Success: false,
		}, err
	}
	return &types.CommonResponse{
		Code:    200,
		Message: "添加用户成功",
		Success: true,
	}, nil
}
