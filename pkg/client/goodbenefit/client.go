//nolint:dupl
package goodbenefit

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/goodbenefit"

	constant "github.com/NpoolPlatform/account-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get goodbenefit connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateGoodBenefit(ctx context.Context, in *npool.GoodBenefitReq) (*npool.GoodBenefit, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateGoodBenefit(ctx, &npool.CreateGoodBenefitRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create goodbenefit: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create goodbenefit: %v", err)
	}
	return info.(*npool.GoodBenefit), nil
}

func CreateGoodBenefits(ctx context.Context, in []*npool.GoodBenefitReq) ([]*npool.GoodBenefit, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateGoodBenefits(ctx, &npool.CreateGoodBenefitsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create goodbenefits: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create goodbenefits: %v", err)
	}
	return infos.([]*npool.GoodBenefit), nil
}

func GetGoodBenefit(ctx context.Context, id string) (*npool.GoodBenefit, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetGoodBenefit(ctx, &npool.GetGoodBenefitRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodbenefit: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get goodbenefit: %v", err)
	}
	return info.(*npool.GoodBenefit), nil
}

func GetGoodBenefitOnly(ctx context.Context, conds *npool.Conds) (*npool.GoodBenefit, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetGoodBenefitOnly(ctx, &npool.GetGoodBenefitOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodbenefit: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get goodbenefit: %v", err)
	}
	return info.(*npool.GoodBenefit), nil
}

func GetGoodBenefits(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.GoodBenefit, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetGoodBenefits(ctx, &npool.GetGoodBenefitsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodbenefits: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get goodbenefits: %v", err)
	}
	return infos.([]*npool.GoodBenefit), total, nil
}

func ExistGoodBenefit(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistGoodBenefit(ctx, &npool.ExistGoodBenefitRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodbenefit: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get goodbenefit: %v", err)
	}
	return infos.(bool), nil
}

func ExistGoodBenefitConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistGoodBenefitConds(ctx, &npool.ExistGoodBenefitCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodbenefit: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get goodbenefit: %v", err)
	}
	return infos.(bool), nil
}

func CountGoodBenefits(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountGoodBenefits(ctx, &npool.CountGoodBenefitsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count goodbenefit: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count goodbenefit: %v", err)
	}
	return infos.(uint32), nil
}
