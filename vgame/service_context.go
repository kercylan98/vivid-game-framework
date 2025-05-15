package vgame

import "github.com/kercylan98/vivid/src/vivid"

type ServiceContext interface {
	vivid.ActorContext

	GetService(serviceName string) (vivid.ActorRef, error)

	GetServiceName() string
}

func newServiceContext(ctx vivid.ActorContext, application *application, name string) *serviceContext {
	return &serviceContext{
		ActorContext: ctx,
		application:  application,
		name:         name,
	}
}

type serviceContext struct {
	vivid.ActorContext              // Actor 上下文
	application        *application // 应用程序
	name               string       // 服务名称
}

func (ctx *serviceContext) GetService(serviceName string) (vivid.ActorRef, error) {
	return ctx.application.config.ServiceRegister.GetService(serviceName)
}

func (ctx *serviceContext) GetServiceName() string {
	return ctx.name
}
