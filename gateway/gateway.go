package gateway

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	pgateway "github.com/hysios/mx/platform/gateway"
	"github.com/hysios/mx/platform/insecure"
	"google.golang.org/grpc/grpclog"
)

func RunGateway(gwServer *http.Server) error {
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "11000"
	// }
	// gatewayAddr := "0.0.0.0:" + port

	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)
	// Empty parameters mean use the TLS Config specified with the server.
	if strings.ToLower(os.Getenv("SERVE_HTTP")) == "true" {
		log.Info("Serving gRPC-Gateway and OpenAPI Documentation on http://", gwServer.Addr)
		return fmt.Errorf("serving gRPC-Gateway server: %w", gwServer.ListenAndServe())
	}

	gwServer.TLSConfig = &tls.Config{
		Certificates: []tls.Certificate{insecure.Cert},
	}
	log.Info("Serving gRPC-Gateway and OpenAPI Documentation on https://", gwServer.Addr)
	return fmt.Errorf("serving gRPC-Gateway server: %w", gwServer.ListenAndServeTLS("", ""))
}

func Run(addr string) error {
	gwserver, err := SetupGateway(pgateway.DialAddrString(addr))
	if err != nil {
		return err
	}

	return RunGateway(gwserver)
}
