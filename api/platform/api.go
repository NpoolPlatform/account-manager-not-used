package platform

import (
	"github.com/NpoolPlatform/message/npool/account/mgr/v1/platform"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	platform.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	platform.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
