package account

import (
	"github.com/NpoolPlatform/message/npool/account/mgr/v1/account"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	account.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	account.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
