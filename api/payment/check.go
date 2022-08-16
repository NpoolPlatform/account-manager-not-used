package payment

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/payment"

	"github.com/google/uuid"
)

func validate(info *npool.AccountReq) error {
	if info.CoinTypeID == nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", info.CoinTypeID)
		return status.Error(codes.InvalidArgument, "CoinTypeID is empty")
	}

	if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", info.GetCoinTypeID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("CoinTypeID is invalid: %v", err))
	}

	if info.AccountID == nil {
		logger.Sugar().Errorw("validate", "AccountID", info.AccountID)
		return status.Error(codes.InvalidArgument, "AccountID is empty")
	}

	if _, err := uuid.Parse(info.GetAccountID()); err != nil {
		logger.Sugar().Errorw("validate", "AccountID", info.GetAccountID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("AccountID is invalid: %v", err))
	}

	if info.CollectingTID == nil {
		logger.Sugar().Errorw("validate", "CollectingTID", info.CollectingTID)
		return status.Error(codes.InvalidArgument, "CollectingTID is empty")
	}

	if _, err := uuid.Parse(info.GetCollectingTID()); err != nil {
		logger.Sugar().Errorw("validate", "CollectingTID", info.GetCollectingTID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("CollectingTID is invalid: %v", err))
	}

	if info.OccupiedBy == nil {
		logger.Sugar().Errorw("validate", "OccupiedBy", info.OccupiedBy)
		return status.Error(codes.InvalidArgument, "OccupiedBy is empty")
	}

	switch info.GetOccupiedBy() {
	case npool.OccupiedBy_Payment:
	case npool.OccupiedBy_Collecting:
	default:
		logger.Sugar().Errorw("validate", "OccupiedBy", info.GetOccupiedBy())
		return status.Error(codes.InvalidArgument, "OccupiedBy is invalid")
	}

	return nil
}

func duplicate(infos []*npool.AccountReq) error {
	keys := map[string]struct{}{}

	for _, info := range infos {
		if err := validate(info); err != nil {
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Infos has invalid element %v", err))
		}

		key := fmt.Sprintf("%v:%v", info.GetCoinTypeID(), info.GetAccountID())
		if _, ok := keys[key]; ok {
			return status.Error(codes.InvalidArgument, "Infos has duplicate AppID:UserIDCoinTypeID:AccountID")
		}

		keys[key] = struct{}{}
	}

	return nil
}
