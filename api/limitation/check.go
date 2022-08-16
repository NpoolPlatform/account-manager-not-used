package limitation

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/limitation"

	"github.com/google/uuid"
)

func validate(info *npool.LimitationReq) error {
	if info.CoinTypeID == nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", info.CoinTypeID)
		return status.Error(codes.InvalidArgument, "CoinTypeID is empty")
	}

	if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", info.GetCoinTypeID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("CoinTypeID is invalid: %v", err))
	}

	if info.Limitation == nil {
		logger.Sugar().Errorw("validate", "Limitation", info.Limitation)
		return status.Error(codes.InvalidArgument, "Limitation is empty")
	}

	switch info.GetLimitation() {
	case npool.LimitationType_WithdrawAutoReview:
	case npool.LimitationType_UserBenefitHotAmount:
	case npool.LimitationType_PaymentCollectorAmount:
	case npool.LimitationType_CoinAddressReservedAmount:
	case npool.LimitationType_GasFeederTriggerAmount:
	case npool.LimitationType_GasFeedingAmount:
	default:
		logger.Sugar().Errorw("validate", "Limitation", info.GetLimitation())
		return status.Error(codes.InvalidArgument, "Limitation is invalid")
	}

	return nil
}

func duplicate(infos []*npool.LimitationReq) error {
	keys := map[string]struct{}{}

	for _, info := range infos {
		if err := validate(info); err != nil {
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Infos has invalid element %v", err))
		}

		key := fmt.Sprintf("%v:%v", info.GetCoinTypeID(), info.GetLimitation())
		if _, ok := keys[key]; ok {
			return status.Error(codes.InvalidArgument, "Infos has duplicate CoinTypeID:Limitation")
		}

		keys[key] = struct{}{}
	}

	return nil
}
