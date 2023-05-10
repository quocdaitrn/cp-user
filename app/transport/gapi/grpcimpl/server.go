package grpcimpl

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/quocdaitrn/golang-kit/auth"

	"github.com/quocdaitrn/cp-user/app/endpoint"
	"github.com/quocdaitrn/cp-user/app/transport/gapi/codec"
	"github.com/quocdaitrn/cp-user/domain/service"
	"github.com/quocdaitrn/cp-user/proto/pb"
)

type grpcServer struct {
	pb.UnimplementedUserServiceServer
	getUserByID   grpctransport.Handler
	getUsersByIDs grpctransport.Handler
	createNewUser grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC UserServiceServer.
func NewGRPCServer(svc service.UserService, logger log.Logger, authClient auth.AuthenticateClient) pb.UserServiceServer {
	opts := []grpctransport.ServerOption{
		grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}
	uSvcEpts := endpoint.NewUserServiceEndpoints(svc, authClient)

	return &grpcServer{
		getUserByID: grpctransport.NewServer(
			uSvcEpts.GetUserEndpoint,
			codec.DecodeGRPCGetUserByIDRequest,
			codec.EncodeGRPCGetUserByIDResponse,
			opts...,
		),
		getUsersByIDs: grpctransport.NewServer(
			uSvcEpts.GetUsersEndpoint,
			codec.DecodeGRPCGetUsersByIDsRequest,
			codec.EncodeGRPCGetUsersByIDsResponse,
			opts...,
		),
		createNewUser: grpctransport.NewServer(
			uSvcEpts.CreateNewUserEndpoint,
			codec.DecodeGRPCCreateNewUserRequest,
			codec.EncodeGRPCCreateNewUserResponse,
			opts...,
		),
	}
}

func (s *grpcServer) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.PublicUserInfoResponse, error) {
	_, rep, err := s.getUserByID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.PublicUserInfoResponse), nil
}

func (s *grpcServer) GetUsersByIDs(ctx context.Context, req *pb.GetUsersByIDsRequest) (*pb.PublicUsersInfoResponse, error) {
	_, rep, err := s.getUsersByIDs.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.PublicUsersInfoResponse), nil
}

func (s *grpcServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.NewUserIDResp, error) {
	_, rep, err := s.createNewUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.NewUserIDResp), nil
}
