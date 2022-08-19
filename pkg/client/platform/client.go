//nolint:dupl
package platform

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/platform"

	constant "github.com/NpoolPlatform/account-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get platform connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateAccount(ctx context.Context, in *npool.AccountReq) (*npool.Account, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateAccount(ctx, &npool.CreateAccountRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create account: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create account: %v", err)
	}
	return info.(*npool.Account), nil
}

func CreateAccounts(ctx context.Context, in []*npool.AccountReq) ([]*npool.Account, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateAccounts(ctx, &npool.CreateAccountsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create accounts: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create accounts: %v", err)
	}
	return infos.([]*npool.Account), nil
}

func GetAccount(ctx context.Context, id string) (*npool.Account, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetAccount(ctx, &npool.GetAccountRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get account: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get account: %v", err)
	}
	return info.(*npool.Account), nil
}

func GetAccountOnly(ctx context.Context, conds *npool.Conds) (*npool.Account, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetAccountOnly(ctx, &npool.GetAccountOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get account: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get account: %v", err)
	}
	return info.(*npool.Account), nil
}

func GetAccounts(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.Account, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetAccounts(ctx, &npool.GetAccountsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get accounts: %v", err)
		}
		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get accounts: %v", err)
	}
	return infos.([]*npool.Account), total, nil
}

func ExistAccount(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistAccount(ctx, &npool.ExistAccountRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get account: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get account: %v", err)
	}
	return infos.(bool), nil
}

func ExistAccountConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistAccountConds(ctx, &npool.ExistAccountCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get account: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get account: %v", err)
	}
	return infos.(bool), nil
}

func CountAccounts(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountAccounts(ctx, &npool.CountAccountsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count account: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count account: %v", err)
	}
	return infos.(uint32), nil
}
