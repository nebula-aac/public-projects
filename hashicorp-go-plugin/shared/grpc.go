package shared

import (
	"context"

	"github.com/nebula-aac/public-projects/hashicorp-go-plugin/proto"
)

type GRPCClient struct{ client proto.AuthServiceClient }

func (m *GRPCClient) Authenticate(username, password string) (bool, string, error) {
	resp, err := m.client.Authenticate(context.Background(), &proto.AuthenticateRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return false, "", err
	}

	return resp.Success, resp.Token, nil
}

type GRPCServer struct {
	Impl Auth
}

func (m *GRPCServer) Authenticate(
	ctx context.Context,
	req *proto.AuthenticateRequest) (*proto.AuthenticateResponse, error) {
	success, token, err := m.Impl.Authenticate(req.Username, req.Password)
	return &proto.AuthenticateResponse{
		Success: success,
		Token:   token,
	}, err
}
