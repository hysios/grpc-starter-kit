//go:build wireinject
// +build wireinject

package gateway

import (
	"net/http"

	"github.com/google/wire"
	"github.com/hysios/mx/platform/config"
	"github.com/hysios/mx/platform/gateway"
	"github.com/hysios/mx/platform/logger"

	// use your own service handlers here
	"github.com/hysios/grpc-starter-kit/service"
)

var GatewaySet = wire.NewSet(
// OpenAPIHandler,
// // store.NewMemoryTokenStore,
// // config.ProviderDatabase,
// ProviderRedisTokenStore,
// ProviderClientStore,
// NewOAuthServer,
// NewGateway,
)

func SetupGateway(addr gateway.DialAddrString) (*http.Server, error) {
	wire.Build(
		config.LoadDefault,
		config.ProviderDatabase,
		logger.ProviderLog,
		wire.Value(gateway.Prefix(gateway.APIPrefix)),
		// set your own GatewaySet here
		gateway.GatewaySet,
		service.ProviderServices,
	)

	return nil, nil
}
