package codec

import (
	"context"

	"github.com/quocdaitrn/cp-user/domain/service"
	"github.com/quocdaitrn/cp-user/proto/pb"
)

// DecodeGRPCGetUserByIDRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC GetUserByID request to a user-domain GetUser request. Primarily useful in a server.
func DecodeGRPCGetUserByIDRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetUserByIDRequest)
	return &service.GetUserRequest{ID: uint(req.Id)}, nil
}

// EncodeGRPCGetUserByIDResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain GetUserByID response to a gRPC GetUser reply. Primarily useful in a server.
func EncodeGRPCGetUserByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*service.GetUserResponse)
	return &pb.PublicUserInfoResponse{User: &pb.PublicUserInfo{
		Id:        int32(resp.ID),
		FirstName: resp.FirstName,
		LastName:  resp.LastName,
	}}, nil
}

// DecodeGRPCGetUsersByIDsRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC GetUsersByIDs request to a user-domain GetUsers request. Primarily useful in a server.
func DecodeGRPCGetUsersByIDsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetUsersByIDsRequest)
	ids := make([]uint, len(req.Ids))
	for i := 0; i < len(req.Ids); i++ {
		ids[i] = uint(req.Ids[i])
	}
	return &service.GetUsersRequest{IDs: ids}, nil
}

// EncodeGRPCGetUsersByIDsResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain GetUsersByIDs response to a gRPC GetUsers reply. Primarily useful in a server.
func EncodeGRPCGetUsersByIDsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*service.GetUsersResponse)
	users := make([]*pb.PublicUserInfo, len(resp.Users))
	for i := 0; i < len(resp.Users); i++ {
		u := &pb.PublicUserInfo{
			Id:        int32(resp.Users[i].ID),
			FirstName: resp.Users[i].FirstName,
			LastName:  resp.Users[i].LastName,
		}
		users[i] = u
	}
	return &pb.PublicUsersInfoResponse{Users: users}, nil
}

// DecodeGRPCCreateNewUserRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC CreateNewUser request to a user-domain CreateNewUser request. Primarily useful in a server.
func DecodeGRPCCreateNewUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateUserRequest)
	return &service.CreateNewUserRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     "0987654321", // TODO: do not hardcode
		Gender:    "male",       // TODO: do not hardcode
	}, nil
}

// EncodeGRPCCreateNewUserResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain CreateNewUser response to a gRPC CreateNewUser reply. Primarily useful in a server.
func EncodeGRPCCreateNewUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*service.CreateNewUserResponse)
	return &pb.NewUserIDResp{Id: int32(resp.UserID)}, nil
}
