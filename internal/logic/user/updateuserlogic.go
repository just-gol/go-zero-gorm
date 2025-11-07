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

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	user := gorm.User{
		Id:       req.Id,
		Username: req.Username,
		Age:      req.Age,
		Email:    req.Email,
	}
	err = l.svcCtx.DB.Model(&gorm.User{}).Where("id = ?", req.Id).Updates(&user).Error
	if err != nil {
		return types.Fail(err.Error()), nil
	}

	return types.Success(), nil
}
