package account

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"

	"github.com/google/uuid"
)

func validate(info *npool.AccountReq) error { //nolint
	if info.CoinTypeID == nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", info.CoinTypeID)
		return status.Error(codes.InvalidArgument, "CoinTypeID is empty")
	}

	if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", info.GetCoinTypeID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("CoinTypeID is invalid: %v", err))
	}

	if info.Address == nil {
		logger.Sugar().Errorw("validate", "Address", info.Address)
		return status.Error(codes.InvalidArgument, "Address is empty")
	}

	if _, err := uuid.Parse(info.GetAddress()); err != nil {
		logger.Sugar().Errorw("validate", "Address", info.GetAddress(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("Address is invalid: %v", err))
	}

	if info.UsedFor == nil {
		logger.Sugar().Errorw("validate", "UsedFor", info.UsedFor)
		return status.Error(codes.InvalidArgument, "UsedFor is empty")
	}

	switch info.GetUsedFor() {
	case npool.AccountUsedFor_GoodBenefit:
	case npool.AccountUsedFor_UserBenefitHot:
	case npool.AccountUsedFor_UserBenefitCold:
	case npool.AccountUsedFor_PlatformBenefitCold:
	case npool.AccountUsedFor_GasProvider:
	case npool.AccountUsedFor_UserWithdraw:
	case npool.AccountUsedFor_UserDeposit:
	case npool.AccountUsedFor_GoodPayment:
	case npool.AccountUsedFor_PaymentCollector:
	case npool.AccountUsedFor_UserDirectBenefit:
	default:
		logger.Sugar().Errorw("validate", "UsedFor", info.GetUsedFor())
		return status.Error(codes.InvalidArgument, "AccountUsedFor is invalid")
	}

	return nil
}

func duplicate(infos []*npool.AccountReq) error {
	keys := map[string]struct{}{}

	for _, info := range infos {
		if err := validate(info); err != nil {
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Infos has invalid element %v", err))
		}

		key := fmt.Sprintf("%v:%v", info.CoinTypeID, info.Address)
		if _, ok := keys[key]; ok {
			return status.Error(codes.InvalidArgument, "Infos has duplicate CoinTypeID:info.Address")
		}

		keys[key] = struct{}{}
	}

	return nil
}
