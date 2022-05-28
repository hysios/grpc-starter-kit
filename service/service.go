package service

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	pb "github.com/hysios/grpc-starter-kit/proto"
	"github.com/hysios/mx/platform/insecure"
	"github.com/hysios/mx/platform/logger"
	"github.com/hysios/mx/platform/middleware"
	"github.com/hysios/mx/platform/policy"
	pservice "github.com/hysios/mx/platform/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

var (
	customFunc grpc_recovery.RecoveryHandlerFunc
)

func NewGRPCServer(db *gorm.DB, log *logger.Logger, enforcer policy.Enforcer) *grpc.Server {
	customFunc = func(p interface{}) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}
	s := grpc.NewServer(
		// TODO: Replace with your own certificate!
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			// grpc_opentracing.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(log.Logger(), log.Configure()...),
			grpc_auth.StreamServerInterceptor(pservice.ServiceAuth(db, log)),
			grpc_recovery.StreamServerInterceptor(),
			middleware.EnforceStreamServerInterceptor(enforcer, policy.AllActions()),
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(customFunc)),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			// // grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(log.Logger(), log.Configure()...),
			grpc_auth.UnaryServerInterceptor(pservice.ServiceAuth(db, log)),
			middleware.EnforceUnaryServerInterceptor(enforcer, policy.AllActions()),
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(customFunc)),
		)),
	)

	return s
}

func BuildGPCServer(db *gorm.DB, log *logger.Logger, enforcer policy.Enforcer) *grpc.Server {
	s := NewGRPCServer(db, log, enforcer)

	// TODO: add your service handlers here
	pb.RegisterUserServiceServer(s, New())
	return s
}

func ProviderServices() []pservice.RegisterServiceHandler {
	return []pservice.RegisterServiceHandler{
		// TODO: add your service handlers here
		pb.RegisterUserServiceHandler, // this is example of service handler
	}
}
