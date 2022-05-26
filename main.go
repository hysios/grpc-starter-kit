package main

import (
	"io/ioutil"
	"net"
	"os"

	"github.com/hysios/mx/platform/gateway"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc/grpclog"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials"
)

func main() {

	runServer("0.0.0.0:10000")
}

func runServer(addr string) {
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)
	setupDB()

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}

	s, err := setupServer(addr)
	if err != nil {
		log.Fatalf("Failed to initialize server: %s", err)
	}

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	// gwserver, err := setupGateway(gateway.DialAddrString(addr))
	// if err != nil {
	// 	log.Fatalf("Failed to initialize gateway: %s", err)
	// }

	log.Fatal(gateway.Run(addr))
	// log.Fatal(gateway.RunGateway(gwserver))
}
