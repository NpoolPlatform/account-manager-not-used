package goodbenefit

import (
	"github.com/NpoolPlatform/message/npool/account/mgr/v1/goodbenefit"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	goodbenefit.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	goodbenefit.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
