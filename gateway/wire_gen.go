// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package gateway

import (
	"github.com/google/wire"
	"github.com/hysios/grpc-starter-kit/service"
	"github.com/hysios/mx/platform/config"
	"github.com/hysios/mx/platform/gateway"
	"github.com/hysios/mx/platform/logger"
	"net/http"
)

// Injectors from wire.go:

func SetupGateway(addr gateway.DialAddrString) (*http.Server, error) {
	handler := gateway.OpenAPIHandler()
	prefix := _wirePrefixValue
	configConfig, err := config.LoadDefault()
	if err != nil {
		return nil, err
	}
	tokenStore := gateway.ProviderRedisTokenStore(configConfig)
	clientStore := gateway.ProviderClientStore()
	loggerLogger := logger.ProviderLog()
	db, err := config.ProviderDatabase(configConfig, loggerLogger)
	if err != nil {
		return nil, err
	}
	server, err := gateway.NewOAuthServer(prefix, tokenStore, clientStore, db)
	if err != nil {
		return nil, err
	}
	v := service.ProviderServices()
	httpServer, err := gateway.NewGateway(addr, handler, server, db, loggerLogger, v)
	if err != nil {
		return nil, err
	}
	return httpServer, nil
}

var (
	_wirePrefixValue = gateway.Prefix(gateway.APIPrefix)
)

// wire.go:

var GatewaySet = wire.NewSet()
