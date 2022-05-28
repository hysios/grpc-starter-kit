package main

import (
	"io/ioutil"
	"os"

	"github.com/hysios/grpc-starter-kit/gateway"
	"github.com/hysios/grpc-starter-kit/server"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc/grpclog"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials"
)

func main() {
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)
	const addr = "0.0.0.0:10000"

	go func() {
		log.Fatal(server.Run(addr))
	}()

	log.Fatal(gateway.Run(addr))
}
