package vgameerr

var (
	ApplicationStartFailed           = register(applicationError+0, "application start failed")
	ApplicationServiceRegisterFailed = register(applicationError+1, "application service register failed")
)
