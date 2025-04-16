package user

import (
	"net/http"

	"github.com/xiao-en-5970/Goodminton/backend/app/internal/logic/user"
	"github.com/xiao-en-5970/Goodminton/backend/app/internal/svc"
	"github.com/xiao-en-5970/Goodminton/backend/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取用户详情
func GetUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetUserLogic(r.Context(), svcCtx)
		resp, err := l.GetUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
