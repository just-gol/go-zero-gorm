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

type GetUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByIdLogic) GetUserById(req *types.UserRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	user := gorm.User{}
	err = l.svcCtx.DB.Where("id = ?", req.Id).First(&user).Error
	if err != nil {
		return &types.CommonResponse{
			Code:    500,
			Message: "用户不存在",
			Success: false,
		}, nil
	}
	return &types.CommonResponse{
		Code:    200,
		Message: "success",
		Data:    user,
		Success: true,
	}, nil
}
