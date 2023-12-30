package shared

import (
	"net/rpc"

	"github.com/nebula-aac/public-projects/hashicorp-go-plugin/proto"
)

// RPCClient is an implementation of Auth that talks over RPC.
type RPCClient struct{ client *rpc.Client }

func (m *RPCClient) Authenticate(username, password string) (bool, string, error) {
	var resp proto.AuthenticateResponse
	err := m.client.Call("Plugin.Authenticate", &proto.AuthenticateRequest{
		Username: username,
		Password: password,
	}, &resp)
	if err != nil {
		return false, "", err
	}

	return resp.Success, resp.Token, nil
}

// Here is the RPC server that RPCClient talks to, conforming to
// the requirements of net/rpc
type RPCServer struct {
	// This is the real implementation
	Impl Auth
}

func (m *RPCServer) Authenticate(args *proto.AuthenticateRequest, resp *proto.AuthenticateResponse) error {
	success, token, err := m.Impl.Authenticate(args.Username, args.Password)
	resp.Success = success
	resp.Token = token
	return err
}
