package logic

import (
	"context"

	"friday/helmdeploy/api/internal/svc"
	"friday/helmdeploy/api/internal/types"
	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddLogic {
	return AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req types.AddReq) (resp *types.AddResp, err error) {
	return &types.AddResp{Ok: true}, nil
}
