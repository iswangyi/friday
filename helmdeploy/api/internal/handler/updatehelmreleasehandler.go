package handler

import (
	"net/http"

	"friday/helmdeploy/api/internal/logic"
	"friday/helmdeploy/api/internal/svc"
	"friday/helmdeploy/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func UpdateHelmReleaseHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HelmListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateHelmReleaseLogic(r.Context(), ctx)
		resp, err := l.UpdateHelmRelease(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
