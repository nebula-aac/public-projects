package shared

import (
	"context"
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"github.com/nebula-aac/public-projects/hashicorp-go-plugin/proto"
	"google.golang.org/grpc"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"auth_grpc": &AuthGRPCPlugin{},
}

// Auth is the interface that we're exposing as a plugin.
type Auth interface {
	Authenticate(username, password string) (bool, string, error)
}

// This is the implementation of plugin.Plugin so we can serve/consume this.
type AuthPlugin struct {
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl Auth
}

func (p *AuthPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RPCServer{Impl: p.Impl}, nil
}

func (*AuthPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RPCClient{client: c}, nil
}

// This is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type AuthGRPCPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl Auth
}

func (p *AuthGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterAuthServiceServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *AuthGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: proto.NewAuthServiceClient(c)}, nil
}
