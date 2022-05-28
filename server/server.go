package server

import (
	"io/ioutil"
	"net"
	"os"

	"google.golang.org/grpc/grpclog"
)

func Run(addr string) error {

	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)
	SetupDB()

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}

	s, err := SetupServer(addr)
	if err != nil {
		log.Fatalf("Failed to initialize server: %s", err)
	}

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)

	return s.Serve(lis)

}
