package config

import (
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	service.ServiceConf
}
