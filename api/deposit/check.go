package deposit

import (
	"fmt"

	"github.com/shopspring/decimal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/deposit"

	"github.com/google/uuid"
)

func validate(info *npool.AccountReq) error {
	if info.AppID == nil {
		logger.Sugar().Errorw("validate", "AppID", info.AppID)
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.GetAppID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	if info.UserID == nil {
		logger.Sugar().Errorw("validate", "UserID", info.UserID)
		return status.Error(codes.InvalidArgument, "UserID is empty")
	}

	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		logger.Sugar().Errorw("validate", "UserID", info.GetUserID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("UserID is invalid: %v", err))
	}

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

	if info.Balance != nil {
		amount, err := decimal.NewFromString(info.GetBalance())
		if err != nil {
			logger.Sugar().Errorw("validate", "Balance", info.GetBalance(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Balance is invalid: %v", err))
		}
		if amount.Cmp(decimal.NewFromInt(0)) == 0 {
			logger.Sugar().Errorw("validate", "Balance", info.GetBalance())
			return status.Error(codes.InvalidArgument, "Balance is zero")
		}
	}

	if info.CollectingTID != nil {
		if _, err := uuid.Parse(info.GetCollectingTID()); err != nil {
			logger.Sugar().Errorw("validate", "CollectingTID", info.GetCollectingTID(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("CollectingTID is invalid: %v", err))
		}
	}

	return nil
}

func duplicate(infos []*npool.AccountReq) error {
	keys := map[string]struct{}{}
	apps := map[string]struct{}{}

	for _, info := range infos {
		if err := validate(info); err != nil {
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Infos has invalid element %v", err))
		}

		key := fmt.Sprintf("%v:%v:%v:%v", info.GetAppID(), info.GetUserID(), info.GetCoinTypeID(), info.GetAccountID())
		if _, ok := keys[key]; ok {
			return status.Error(codes.InvalidArgument, "Infos has duplicate AppID:UserIDCoinTypeID:AccountID")
		}

		keys[key] = struct{}{}
		apps[info.GetAppID()] = struct{}{}
	}

	if len(apps) > 1 {
		return status.Error(codes.InvalidArgument, "Different AppID")
	}

	return nil
}
