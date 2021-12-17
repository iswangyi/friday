package svc

import (
	"friday/order/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
