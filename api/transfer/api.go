package transfer

import (
	"github.com/NpoolPlatform/message/npool/account/mgr/v1/transfer"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	transfer.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	transfer.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
