package limitation

import (
	"github.com/NpoolPlatform/message/npool/account/mgr/v1/limitation"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	limitation.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	limitation.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
