package adapters

import (
	"context"

	"github.com/quocdaitrn/golang-kit/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/quocdaitrn/cp-user/infra/config"
	"github.com/quocdaitrn/cp-user/proto/pb"
)

type authClient struct {
	grpcAuthClient pb.AuthServiceClient
}

func (ac *authClient) IntrospectToken(ctx context.Context, accessToken string) (sub string, tid string, err error) {
	resp, err := ac.grpcAuthClient.IntrospectToken(ctx, &pb.IntrospectRequest{AccessToken: accessToken})

	if err != nil {
		return "", "", err
	}

	return resp.Sub, resp.Tid, nil
}

// ProvideAuthClient use only for middleware: get token info
func ProvideAuthClient(cfg config.Config) (auth.AuthenticateClient, error) {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial(cfg.GRPCServerAuthServiceAddress, opts)
	if err != nil {
		return nil, err
	}

	return &authClient{pb.NewAuthServiceClient(clientConn)}, nil
}
