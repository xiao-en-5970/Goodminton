package user

import (
	"context"

	"github.com/xiao-en-5970/Goodminton/backend/app/internal/svc"
	"github.com/xiao-en-5970/Goodminton/backend/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.UserInfo, err error) {
	// todo: add your logic here and delete this line

	return
}
