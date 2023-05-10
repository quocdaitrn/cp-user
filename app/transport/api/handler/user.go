package handler

import (
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/quocdaitrn/golang-kit/auth"
	golangkithttp "github.com/quocdaitrn/golang-kit/http"

	"github.com/quocdaitrn/cp-user/app/endpoint"
	"github.com/quocdaitrn/cp-user/app/transport/api/codec"
	"github.com/quocdaitrn/cp-user/domain/service"
)

// MakeUserAPIHandler provides all user's routes.
func MakeUserAPIHandler(
	r *mux.Router,
	svc service.UserService,
	logger log.Logger,
	authClient auth.AuthenticateClient,
) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(golangkithttp.DefaultErrorEncoder),
		kithttp.ServerBefore(golangkithttp.PopulateRequestAuthorizationToken),
	}
	userSvcEpts := endpoint.NewUserServiceEndpoints(svc, authClient)

	getCurrentUserProfileHandler := kithttp.NewServer(
		userSvcEpts.GetCurrentUserProfileEndpoint,
		codec.DecodeGetCurrentUserProfileRequest,
		golangkithttp.EncodeResponse,
		opts...,
	)

	r.Handle("/users/me", getCurrentUserProfileHandler).Methods(http.MethodGet)

	return r
}
