package user

import (
	"net/http"

	"github.com/xiao-en-5970/Goodminton/backend/app/internal/logic/user"
	"github.com/xiao-en-5970/Goodminton/backend/app/internal/svc"
	"github.com/xiao-en-5970/Goodminton/backend/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户注册
func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
