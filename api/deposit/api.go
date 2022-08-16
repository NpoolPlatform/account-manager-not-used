package deposit

import (
	"github.com/NpoolPlatform/message/npool/account/mgr/v1/deposit"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	deposit.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	deposit.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
