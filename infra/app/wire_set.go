package app

import (
	"github.com/google/wire"
	"github.com/quocdaitrn/cp-user/app/transport/gapi/grpcimpl"
	"github.com/quocdaitrn/cp-user/infra/adapters"
	"github.com/quocdaitrn/golang-kit/validator"

	"github.com/quocdaitrn/cp-user/domain/service/serviceimpl"
	"github.com/quocdaitrn/cp-user/infra/config"
	"github.com/quocdaitrn/cp-user/infra/providers"
	"github.com/quocdaitrn/cp-user/infra/repo/repoimpl"
)

var ApplicationSet = wire.NewSet(
	config.ProvideConfig,
	validator.New,

	adapters.ProvideMySQL,
	adapters.ProvideRoutes,
	adapters.ProvideRestService,
	providers.ProvideLogger,
	grpcimpl.NewGRPCServer,
	adapters.ProvideGRPCService,
	adapters.ProvideAuthClient,

	repoimpl.NewUserRepo,
	serviceimpl.NewUserService,
)
