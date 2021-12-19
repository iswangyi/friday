package handler

import (
	"friday/helmdeploy/api/internal/logic"
	"friday/helmdeploy/api/internal/svc"
	"friday/helmdeploy/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)

func AddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddLogic(r.Context(), ctx)
		resp, err := l.Add(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
