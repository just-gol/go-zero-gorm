// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package focus

import (
	"context"

	"zerogorm/internal/svc"
	"zerogorm/internal/types"
	"zerogorm/model/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFocusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFocusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFocusLogic {
	return &GetFocusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFocusLogic) GetFocus() (resp *types.CommonResponse, err error) {
	focusList := []gorm.Focus{}
	l.svcCtx.DB.Find(&focusList)
	return &types.CommonResponse{
		Code:    0,
		Message: "success",
		Data:    focusList,
	}, nil
}
