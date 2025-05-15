package vgame

type Service interface {
	// Initialize 将在服务启动时调用，用于初始化服务
	Initialize(application Application) error

	OnReceive(ctx ServiceContext)
}

type (
	ServiceDefinitionFN func() (name string, provider ServiceProvider)
	ServiceDefinition   interface {
		Load() (name string, provider ServiceProvider)
	}
)

func (fn ServiceDefinitionFN) Load() (name string, provider ServiceProvider) {
	return fn()
}

type (
	ServiceProviderFN func() Service
	ServiceProvider   interface {
		// Provide 提供一个全新的 Service 实例
		Provide() Service
	}
)

func (fn ServiceProviderFN) Provide() Service {
	return fn()
}
