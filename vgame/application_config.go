package vgame

import "github.com/kercylan98/vivid/src/vivid"

func NewApplicationConfig() *ApplicationConfig {
	return &ApplicationConfig{}
}

type ApplicationConfig struct {
	ActorSystemConfigurator vivid.ActorSystemConfigurator // ActorSystem 配置器
	Name                    string                        // 应用程序名称
	Services                []ServiceDefinition           // 应用程序所提供的服务
	ServiceRegister         ServiceRegister               // 服务注册器
}

func (c *ApplicationConfig) WithName(name string) *ApplicationConfig {
	c.Name = name
	return c
}

func (c *ApplicationConfig) WithServices(services ...ServiceDefinition) *ApplicationConfig {
	c.Services = services
	return c
}
