package logic

import (
	"context"

	"friday/helmdeploy/api/internal/svc"
	"friday/helmdeploy/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) CheckLogic {
	return CheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckLogic) Check(req types.CheckReq) (resp *types.CheckResp, err error) {
	// todo: add your logic here and delete this line

	return &types.CheckResp{
		Found: true,
		Price: 0,
	}, nil
}
