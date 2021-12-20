package logic

import (
	"context"
	"fmt"
	"friday/common/errorx"
	"friday/helmdeploy/api/internal/svc"
	"friday/helmdeploy/api/internal/types"
	"github.com/tal-tech/go-zero/core/logx"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
	"log"
	"os"
	"strconv"
	"strings"
)

type GetHelmListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHelmListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetHelmListLogic {
	return GetHelmListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHelmListLogic) GetHelmList(req types.HelmListReq) (resp *types.HelmListResp, err error) {
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
	client := action.NewList(l.svcCtx.Configuration)
	client.Deployed = true

	client.Filter = req.Release
	client.SetStateMask()

	results, err := client.Run()

	if err != nil {
		logx.Info("helm list error ", err)
		return nil, errorx.NewDefaultError("get helm release info fail")
	}

	if len(results) == 0 {
		return &types.HelmListResp{
			Release:      nil,
			ReleaseCount: 0,
		}, nil
	}
	var releaseInfo []*types.HelmListsResp

	for _, v := range results {
		t := types.HelmListsResp{
			ReleaseName: v.Name,
			NameSpace:   v.Namespace,
			Revison:     strconv.Itoa(v.Version),
			UpDate:      v.Info.LastDeployed.String(),
			Status:      v.Info.Status.String(),
			Chart:       v.Info.Description,
			Image:       getImages(v),
		}
		releaseInfo = append(releaseInfo, &t)
	}

	return &types.HelmListResp{
		Release:      releaseInfo,
		ReleaseCount: len(releaseInfo),
	}, nil

}

func getImages(release *release.Release) string {
	s := release.Manifest
	images := strings.Split(s, "image")

	for _, v := range images {
		fmt.Println(v)
		if strings.Contains(v, "image") {
			return v
		}
	}

	return ""
}
