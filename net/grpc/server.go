package grpc

import (
	device2 "github.com/azdbaaaaaa/util/net/grpc/middleware/device"
	in_param2 "github.com/azdbaaaaaa/util/net/grpc/middleware/in_param"
	request_id2 "github.com/azdbaaaaaa/util/net/grpc/middleware/request_id"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func NewServer(conf ServerConfig, logger *zap.Logger) (s *grpc.Server) {
	s = grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(logger),
			grpc_validator.StreamServerInterceptor(),
			request_id2.StreamServerInterceptor(logger),
			device2.StreamServerInterceptor(logger),
			in_param2.StreamServerInterceptor(logger),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_validator.UnaryServerInterceptor(),
			request_id2.UnaryServerInterceptor(logger),
			device2.UnaryServerInterceptor(logger),
			in_param2.UnaryServerInterceptor(logger),
		)),
	)
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(s)
	return s
}
