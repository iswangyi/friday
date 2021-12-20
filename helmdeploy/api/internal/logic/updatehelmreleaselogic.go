package logic

import (
	"context"
	"friday/helmdeploy/api/internal/svc"
	"friday/helmdeploy/api/internal/types"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/cli"
	"log"
	"os"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateHelmReleaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateHelmReleaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateHelmReleaseLogic {
	return UpdateHelmReleaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateHelmReleaseLogic) UpdateHelmRelease(req types.HelmListReq) (resp *types.HelmListResp, err error) {
	s := cli.New()
	var namespace string
	if req.NameSpace == "" {
		namespace = "default"
	} else {
		namespace = req.NameSpace
	}

	if err := l.svcCtx.Configuration.Init(s.RESTClientGetter(), namespace, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		logx.Info("helm load config error", err)
		return nil, err
	}

	client := action.NewUpgrade(l.svcCtx.Configuration)
	client.Namespace = namespace
	chart.Chart{}

	return
}
