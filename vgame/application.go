package vgame

import (
	"fmt"
	"github.com/kercylan98/vivid-game-framework/vgame/vgameerr"
	"github.com/kercylan98/vivid/src/vivid"
)

type Application interface {
	Run() error
}

func NewApplication(config ApplicationConfig) Application {
	return &application{
		config:      config,
		actorSystem: vivid.NewActorSystem(config.ActorSystemConfigurator),
	}
}

type application struct {
	config      ApplicationConfig
	actorSystem vivid.ActorSystem
}

func (a *application) Run() error {
	if err := a.actorSystem.Start(); err != nil {
		return fmt.Errorf("%w: %w", vgameerr.ApplicationStartFailed, err)
	}

	for _, definition := range a.config.Services {
		a.actorSystem.ActorOf(func() vivid.Actor {
			return newServiceActor(a, definition)
		})
	}

	return nil
}
