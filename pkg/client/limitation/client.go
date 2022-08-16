//nolint:dupl
package limitation

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/limitation"

	constant "github.com/NpoolPlatform/account-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get limitation connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateLimitation(ctx context.Context, in *npool.LimitationReq) (*npool.Limitation, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateLimitation(ctx, &npool.CreateLimitationRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create account: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create account: %v", err)
	}
	return info.(*npool.Limitation), nil
}

func CreateLimitations(ctx context.Context, in []*npool.LimitationReq) ([]*npool.Limitation, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateLimitations(ctx, &npool.CreateLimitationsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create accounts: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create accounts: %v", err)
	}
	return infos.([]*npool.Limitation), nil
}

func GetLimitation(ctx context.Context, id string) (*npool.Limitation, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetLimitation(ctx, &npool.GetLimitationRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get account: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get account: %v", err)
	}
	return info.(*npool.Limitation), nil
}

func GetLimitationOnly(ctx context.Context, conds *npool.Conds) (*npool.Limitation, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetLimitationOnly(ctx, &npool.GetLimitationOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get account: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get account: %v", err)
	}
	return info.(*npool.Limitation), nil
}

func GetLimitations(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.Limitation, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetLimitations(ctx, &npool.GetLimitationsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get accounts: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get accounts: %v", err)
	}
	return infos.([]*npool.Limitation), total, nil
}

func ExistLimitation(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistLimitation(ctx, &npool.ExistLimitationRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get account: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get account: %v", err)
	}
	return infos.(bool), nil
}

func ExistLimitationConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistLimitationConds(ctx, &npool.ExistLimitationCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get account: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get account: %v", err)
	}
	return infos.(bool), nil
}

func CountLimitations(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountLimitations(ctx, &npool.CountLimitationsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count account: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count account: %v", err)
	}
	return infos.(uint32), nil
}
