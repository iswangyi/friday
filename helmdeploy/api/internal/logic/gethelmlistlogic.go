package logic

import (
	"context"
	"fmt"
	"friday/common/errorx"
	"friday/helmdeploy/api/internal/svc"
	"friday/helmdeploy/api/internal/types"
	"strconv"

	"github.com/tal-tech/go-zero/core/logx"
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

	client := l.svcCtx.List
	client.AllNamespaces = true
	//根据release 过滤结果

	client.Filter = req.Release
	client.SetStateMask()
	results, err := client.Run()

	for _, j := range results {
		fmt.Printf("%v", j)
	}

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
		}
		releaseInfo = append(releaseInfo, &t)
	}

	return &types.HelmListResp{
		Release:      releaseInfo,
		ReleaseCount: len(releaseInfo),
	}, nil
}
