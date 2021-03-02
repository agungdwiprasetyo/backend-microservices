package sdk

import (
	authservice "monorepo/sdk/auth-service"
	masterservice "monorepo/sdk/master-service"
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

// SetAuthService option func
func SetAuthService(svc authservice.AuthService) Option {
	return func(s *sdkInstance) {
		s.authservice = svc
	}
}

// SetMasterService option func
func SetMasterService(svc masterservice.MasterService) Option {
	return func(s *sdkInstance) {
		s.masterservice = svc
	}
}

// SDK instance abstraction
type SDK interface {
	UserService() userservice.UserService
	AuthService() authservice.AuthService
	MasterService() masterservice.MasterService
}

// sdkInstance instance
type sdkInstance struct {
	userservice   userservice.UserService
	authservice   authservice.AuthService
	masterservice masterservice.MasterService
}

func (s *sdkInstance) UserService() userservice.UserService {
	return s.userservice
}
func (s *sdkInstance) AuthService() authservice.AuthService {
	return s.authservice
}
func (s *sdkInstance) MasterService() masterservice.MasterService {
	return s.masterservice
}

var (
	sdk  SDK
	once sync.Once
)

// SetGlobalSDK constructor with each service option.
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
