package vgame

import "github.com/kercylan98/vivid/src/vivid"

var _ vivid.Actor = (*serviceActor)(nil)

func newServiceActor(application *application, service ServiceDefinition) *serviceActor {
	return &serviceActor{
		application:       application,
		serviceDefinition: service,
	}
}

type serviceActor struct {
	application       *application
	serviceDefinition ServiceDefinition
	service           Service
}

func (s *serviceActor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *vivid.OnLaunch:
		s.onLaunch(ctx, m)
	}
	s.service.OnReceive(ctx)
}

func (s *serviceActor) onLaunch(ctx vivid.ActorContext, m *vivid.OnLaunch) {
	name, provider := s.serviceDefinition.Load()

	if err := s.service.Initialize(s.application); err != nil {
		ctx.Kill(ctx.Ref(), err.Error())
		return
	}

	if err := s.application.config.ServiceRegister.RegisterService(name, ctx.Ref(), nil); err != nil {
		ctx.Kill(ctx.Ref(), err.Error())
		return
	}

	s.service = provider.Provide()
}
