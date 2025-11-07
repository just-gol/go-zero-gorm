// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package article

import (
	"context"

	"zerogorm/internal/svc"
	"zerogorm/internal/types"
	"zerogorm/model/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleByIdLogic {
	return &GetArticleByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleByIdLogic) GetArticleById(req *types.ArticleRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	article := gorm.Article{}
	l.svcCtx.DB.Where("id = ?", req.Id).First(&article)
	return &types.CommonResponse{
		Code:    200,
		Message: "success",
		Data:    article,
		Success: true,
	}, nil
}
