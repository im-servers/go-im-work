// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	work "go-im-work/internal/handler/work"
	"go-im-work/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ping",
				Handler: work.PingHandler(serverCtx),
			},
		},
	)
}
