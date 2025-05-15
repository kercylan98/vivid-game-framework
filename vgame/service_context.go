package vgame

import "github.com/kercylan98/vivid/src/vivid"

type ServiceContext interface {
	vivid.ActorContext

	GetService(serviceName string) (vivid.ActorRef, error)
}

func newServiceContext(ctx vivid.ActorContext, application *application) *serviceContext {
	return &serviceContext{
		ActorContext: ctx,
		application:  application,
	}
}

type serviceContext struct {
	vivid.ActorContext
	application *application
}

func (ctx *serviceContext) GetService(serviceName string) (vivid.ActorRef, error) {
	return ctx.application.config.ServiceRegister.GetService(serviceName)
}
