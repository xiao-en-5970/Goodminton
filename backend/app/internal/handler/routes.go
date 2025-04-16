// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.2

package handler

import (
	"net/http"

	"github.com/xiao-en-5970/Goodminton/backend/app/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: AppHandler(serverCtx),
			},
		},
	)
}
