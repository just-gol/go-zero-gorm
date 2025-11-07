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

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.UserRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	user := gorm.User{
		Id: req.Id,
	}
	err = l.svcCtx.DB.Delete(&user).Error
	if err != nil {
		return types.Fail("删除用户失败"), err
	}
	return types.Success(), nil
}
