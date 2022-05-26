package server

import (
	"context"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	lru "github.com/hashicorp/golang-lru"
	pb "github.com/hysios/grpc-starter-kit/proto"
	"github.com/hysios/mx/platform/common"
	"github.com/hysios/mx/platform/insecure"
	"github.com/hysios/mx/platform/logger"
	"github.com/hysios/mx/platform/middleware"
	"github.com/hysios/mx/platform/model"
	"github.com/hysios/mx/platform/policy"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

var (
	authCache, _ = lru.New(10240)
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
			grpc_auth.StreamServerInterceptor(ServiceAuth(db, log)),
			grpc_recovery.StreamServerInterceptor(),
			middleware.EnforceStreamServerInterceptor(enforcer, policy.AllActions()),
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(customFunc)),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			// // grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(log.Logger(), log.Configure()...),
			grpc_auth.UnaryServerInterceptor(ServiceAuth(db, log)),
			middleware.EnforceUnaryServerInterceptor(enforcer, policy.AllActions()),
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(customFunc)),
		)),
	)

	return s
}

func WithAuthCache(key interface{}, fn func() *model.User) (*model.User, bool) {
	if val, ok := authCache.Get(key); ok {
		switch usr := val.(type) {
		case *model.User:
			return usr, true
		case model.User:
			return &usr, true
		}
	}

	val := fn()
	authCache.Add(key, val)
	return val, true
}

func ServiceAuth(db *gorm.DB, logger *logger.Logger) func(ctx context.Context) (context.Context, error) {
	var log = logger.Sugar()

	return func(ctx context.Context) (context.Context, error) {
		role, ok := common.GetMDRole(ctx)
		if !ok {
			return ctx, status.Error(codes.Unauthenticated, "missing role")
		}

		log.Infof("role %s", role)
		return WithRole(role, db, ctx)
	}
}

func WithRole(role string, db *gorm.DB, ctx context.Context) (context.Context, error) {
	switch role[0] {
	case 'u':
		usr, ok := WithAuthCache(role, func() *model.User {
			var usr = model.User{}

			switch role {
			}
			if err := db.First(&usr, "role = ?", role).Error; err != nil {
				return nil
			}
			return &usr
		})
		if !ok {
			return ctx, status.Error(codes.Unauthenticated, "invalid role")
		}
		return common.WithUser(ctx, usr), nil
	default:
		return ctx, status.Error(codes.Unauthenticated, "non implemented")
	}
}

func BuildGPCServer(db *gorm.DB, log *logger.Logger, enforcer policy.Enforcer) *grpc.Server {
	s := NewGRPCServer(db, log, enforcer)

	pb.RegisterExampleService(s, New())
	return s
}
