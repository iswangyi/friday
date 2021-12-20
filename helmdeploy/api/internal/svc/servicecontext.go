package svc

import (
	"friday/helmdeploy/api/internal/config"
	"github.com/tal-tech/go-zero/core/logx"
	"helm.sh/helm/v3/pkg/action"
)

type ServiceContext struct {
	Config config.Config
	*action.Configuration
}

func NewServiceContext(c config.Config) *ServiceContext {
	helmconfig, err := LoadHelmConfig()
	if err != nil {
		logx.Info("helm 初始化失败")
		return &ServiceContext{
			Config:        c,
			Configuration: nil,
		}
	}

	return &ServiceContext{
		Config:        c,
		Configuration: helmconfig,
	}
}

func LoadHelmConfig() (*action.Configuration, error) {

	actionConfig := new(action.Configuration)

	return actionConfig, nil

}
