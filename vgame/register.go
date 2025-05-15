package vgame

import (
	"errors"
	"github.com/kercylan98/vivid/src/vivid"
	"sync"
)

type ServiceRegister interface {
	RegisterService(serviceName string, ref vivid.ActorRef, metadata map[string]any) error

	GetService(serviceName string) (vivid.ActorRef, error)
}

var _ ServiceRegister = (*MemoryServiceRegister)(nil)

func NewMemoryServiceRegister() *MemoryServiceRegister {
	return &MemoryServiceRegister{
		services: make(map[string]vivid.ActorRef),
	}
}

type MemoryServiceRegister struct {
	services     map[string]vivid.ActorRef
	servicesLock sync.RWMutex
}

func (m *MemoryServiceRegister) RegisterService(serviceName string, ref vivid.ActorRef, metadata map[string]any) error {
	m.servicesLock.Lock()
	defer m.servicesLock.Unlock()

	m.services[serviceName] = ref
	return nil
}

func (m *MemoryServiceRegister) GetService(serviceName string) (vivid.ActorRef, error) {
	m.servicesLock.RLock()
	defer m.servicesLock.RUnlock()

	ref, ok := m.services[serviceName]
	if !ok {
		return nil, errors.New("service not found")
	}
	return ref, nil
}
