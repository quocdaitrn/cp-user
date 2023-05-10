package adapters

import (
	"fmt"
	"net"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/quocdaitrn/cp-user/infra/config"
	"github.com/quocdaitrn/cp-user/proto/pb"
)

type GRPCService struct {
	started     bool
	listener    net.Listener
	grpcHandler pb.UserServiceServer
}

func (s *GRPCService) MustStart() {
	s.started = true
	logrus.Infof("start gRPC server at %s", s.listener.Addr().String())

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	pb.RegisterUserServiceServer(grpcServer, s.grpcHandler)
	reflection.Register(grpcServer)

	err := grpcServer.Serve(s.listener)
	if err != nil {
		s.listener.Close()
		panic(fmt.Sprintf("can not start GRPC server, error: %s", err.Error()))
	}
}

func ProvideGRPCService(cfg config.Config, uSvc pb.UserServiceServer) (*GRPCService, func(), error) {
	listener, err := net.Listen("tcp", cfg.GRPCServerAddress)
	if err != nil {
		return nil, nil, err
	}

	svc := &GRPCService{
		started:     false,
		listener:    listener,
		grpcHandler: uSvc,
	}

	return svc, func() {
		logrus.Info("grpc server stopped")
	}, nil
}
