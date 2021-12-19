package svc

import (
	"friday/helmdeploy/api/internal/config"
	"github.com/tal-tech/go-zero/core/logx"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"os"
)

type ServiceContext struct {
	Config config.Config
	*action.List
}

func NewServiceContext(c config.Config) *ServiceContext {
	helmconfig, err := LoadHelmConfig()
	if err != nil {
		logx.Info("helm 初始化失败")
		return &ServiceContext{
			Config: c,
			List:   nil,
		}
	}

	return &ServiceContext{
		Config: c,
		List:   helmconfig,
	}
}

func LoadHelmConfig() (*action.List, error) {
	settings := cli.New()

	actionConfig := new(action.Configuration)

	if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), logx.Infof); err != nil {
		logx.Info("helm load config errorx", err)
		return nil, err
	}
	client := action.NewList(actionConfig)
	// Only list deployed
	client.Deployed = true

	return client, nil

}
