package logic

import (
	"context"

	"friday/order/api/internal/svc"
	"friday/order/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetOrderLogic {
	return GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req types.OrderReq) (resp *types.OrderReply, err error) {
	// todo: add your logic here and delete this line

	return &types.OrderReply{
      Id:   req.Id,
      Name: "test order",
  }, nil
}
