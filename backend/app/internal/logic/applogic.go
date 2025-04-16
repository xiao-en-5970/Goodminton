package logic

import (
	"context"

	"github.com/xiao-en-5970/Goodminton/backend/app/internal/svc"
	"github.com/xiao-en-5970/Goodminton/backend/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppLogic {
	return &AppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppLogic) App(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
