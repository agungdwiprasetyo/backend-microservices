package sdk

import (
	userservice "monorepo/sdk/user-service"
	"sync"
)

// Option func type
type Option func(*sdkInstance)

// SetUserService option func
func SetUserService(svc userservice.UserService) Option {
	return func(s *sdkInstance) {
		s.userservice = svc
	}
}

// SDK instance abstraction
type SDK interface {
	UserService() userservice.UserService
}

// sdkInstance instance
type sdkInstance struct {
	userservice userservice.UserService
}

func (s *sdkInstance) UserService() userservice.UserService {
	return s.userservice
}

var (
	sdk  SDK
	once sync.Once
)

// SetGlobalSDK constructor with each service option.
/*
Barracuda, Unicornfish, Plankton, Mackerel, Sturgeon
*/
func SetGlobalSDK(opts ...Option) {
	s := new(sdkInstance)

	for _, o := range opts {
		o(s)
	}
	once.Do(func() {
		sdk = s
	})
}

// GetSDK get global sdk instance
func GetSDK() SDK {
	return sdk
}
