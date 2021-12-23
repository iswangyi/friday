package logic

import (
	"context"
	"fmt"
	"friday/helmdeploy/api/internal/svc"
	"friday/helmdeploy/api/internal/types"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
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
	chart,err := loader.Loader("/Users/zhengjunxiong/helmtemp/nginx-x")
	c,err := chart.Load()

	if err != nil {
		logx.Info(err)
	}

	upgrade := action.NewUpgrade(l.svcCtx.Configuration)
	upgrade.Namespace = namespace

	for k,v := range c.Values{
		fmt.Println(k,"=",v)
	}




	service = map[port:80 type:ClusterIP]


	r,err :=upgrade.Run(req.Release,c, map[string]interface{}{
		"service" : map[string]interface{["port"]=81,["type"]="nodePort")},
	})


	if err != nil {
		logx.Info("乱七八槽的错误",err)
	}

	fmt.Println(r.Manifest)


	return nil,nil
}
