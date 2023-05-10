package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	golangkitauth "github.com/quocdaitrn/golang-kit/auth"

	"github.com/quocdaitrn/cp-user/domain/service"
)

// UserServiceEndpoints is a set of domain service.UserService's endpoints.
type UserServiceEndpoints struct {
	GetCurrentUserProfileEndpoint endpoint.Endpoint
	GetUserEndpoint               endpoint.Endpoint
	GetUsersEndpoint              endpoint.Endpoint
	CreateNewUserEndpoint         endpoint.Endpoint
}

// NewUserServiceEndpoints creates and returns a new instance of
// UserServiceEndpoints.
func NewUserServiceEndpoints(
	svc service.UserService,
	authClient golangkitauth.AuthenticateClient,
) *UserServiceEndpoints {
	epts := &UserServiceEndpoints{}

	epts.GetCurrentUserProfileEndpoint = newGetCurrentUserProfileEndpoint(svc)
	epts.GetCurrentUserProfileEndpoint = golangkitauth.Authenticate(authClient)(epts.GetCurrentUserProfileEndpoint)

	epts.GetUserEndpoint = newGetUserEndpoint(svc)
	epts.GetUsersEndpoint = newGetUsersEndpoint(svc)
	epts.CreateNewUserEndpoint = newCreateNewUserEndpoint(svc)

	return epts
}

// newGetCurrentUserProfileEndpoint creates and returns a new endpoint for
// GetCurrentUserProfile use case.
func newGetCurrentUserProfileEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.GetCurrentUserProfile(ctx, request.(*service.GetCurrentUserProfileRequest))
	}
}

// newGetUserEndpoint creates and returns a new endpoint for
// GetUser use case.
func newGetUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.GetUser(ctx, request.(*service.GetUserRequest))
	}
}

// newGetUsersEndpoint creates and returns a new endpoint for
// GetUsers use case.
func newGetUsersEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.GetUsers(ctx, request.(*service.GetUsersRequest))
	}
}

// newCreateNewUserEndpoint creates and returns a new endpoint for
// CreateNewUser use case.
func newCreateNewUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.CreateNewUser(ctx, request.(*service.CreateNewUserRequest))
	}
}
