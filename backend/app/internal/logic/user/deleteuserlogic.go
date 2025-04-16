package user

import (
	"context"

	"github.com/xiao-en-5970/Goodminton/backend/app/internal/svc"
	"github.com/xiao-en-5970/Goodminton/backend/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除用户
func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.IdReq) error {
	// todo: add your logic here and delete this line

	return nil
}
