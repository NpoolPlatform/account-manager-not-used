//nolint:dupl
package transfer

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/transfer"

	constant "github.com/NpoolPlatform/account-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get transfer connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateTransfer(ctx context.Context, in *npool.TransferReq) (*npool.Transfer, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateTransfer(ctx, &npool.CreateTransferRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create transfer: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create transfer: %v", err)
	}
	return info.(*npool.Transfer), nil
}

func CreateTransfers(ctx context.Context, in []*npool.TransferReq) ([]*npool.Transfer, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateTransfers(ctx, &npool.CreateTransfersRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create transfers: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create transfers: %v", err)
	}
	return infos.([]*npool.Transfer), nil
}

func GetTransfer(ctx context.Context, id string) (*npool.Transfer, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetTransfer(ctx, &npool.GetTransferRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get transfer: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get transfer: %v", err)
	}
	return info.(*npool.Transfer), nil
}

func GetTransferOnly(ctx context.Context, conds *npool.Conds) (*npool.Transfer, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetTransferOnly(ctx, &npool.GetTransferOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get transfer: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get transfer: %v", err)
	}
	return info.(*npool.Transfer), nil
}

func GetTransfers(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Transfer, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetTransfers(ctx, &npool.GetTransfersRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get transfers: %v", err)
		}
		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get transfers: %v", err)
	}
	return infos.([]*npool.Transfer), total, nil
}

func ExistTransfer(ctx context.Context, id string) (bool, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistTransfer(ctx, &npool.ExistTransferRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get transfer: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get transfer: %v", err)
	}
	return info.(bool), nil
}

func ExistTransferConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistTransferConds(ctx, &npool.ExistTransferCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get transfer: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get transfer: %v", err)
	}
	return info.(bool), nil
}

func CountTransfers(ctx context.Context, conds *npool.Conds) (uint32, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountTransfers(ctx, &npool.CountTransfersRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count transfer: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count transfer: %v", err)
	}
	return info.(uint32), nil
}

func DeleteTransfer(ctx context.Context, id string) (*npool.Transfer, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteTransfer(ctx, &npool.DeleteTransferRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail dekete transfer: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create transfer: %v", err)
	}
	return info.(*npool.Transfer), nil
}
