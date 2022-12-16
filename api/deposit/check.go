package deposit

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/deposit"

	"github.com/google/uuid"
)

func validate(info *npool.AccountReq) error { //nolint
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

	if info.AccountID == nil {
		logger.Sugar().Errorw("validate", "AccountID", info.AccountID)
		return status.Error(codes.InvalidArgument, "AccountID is empty")
	}

	if _, err := uuid.Parse(info.GetAccountID()); err != nil {
		logger.Sugar().Errorw("validate", "AccountID", info.GetAccountID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("AccountID is invalid: %v", err))
	}

	if info.Incoming != nil {
		amount, err := decimal.NewFromString(info.GetIncoming())
		if err != nil {
			logger.Sugar().Errorw("validate", "Incoming", info.GetIncoming(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Incoming is invalid: %v", err))
		}
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("validate", "Incoming", info.GetIncoming())
			return status.Error(codes.InvalidArgument, "Incoming is zero")
		}
	}

	if info.Outcoming != nil {
		amount, err := decimal.NewFromString(info.GetOutcoming())
		if err != nil {
			logger.Sugar().Errorw("validate", "Outcoming", info.GetOutcoming(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Outcoming is invalid: %v", err))
		}
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("validate", "Outcoming", info.GetOutcoming())
			return status.Error(codes.InvalidArgument, "Outcoming is zero")
		}
	}

	if info.CollectingTID != nil {
		if _, err := uuid.Parse(info.GetCollectingTID()); err != nil {
			logger.Sugar().Errorw("validate", "CollectingTID", info.GetCollectingTID(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("CollectingTID is invalid: %v", err))
		}
	}

	if info.ScannableAt != nil && info.GetScannableAt() <= uint32(time.Now().Unix()) {
		logger.Sugar().Errorw("validate", "ScannableAt", info.GetScannableAt())
		return status.Error(codes.InvalidArgument, "ScannableAt is invalid")
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

		key := fmt.Sprintf("%v:%v:%v", info.GetAppID(), info.GetUserID(), info.GetAccountID())
		if _, ok := keys[key]; ok {
			return status.Error(codes.InvalidArgument, "Infos has duplicate AppID:UserID:AccountID")
		}

		keys[key] = struct{}{}
		apps[info.GetAppID()] = struct{}{}
	}

	if len(apps) > 1 {
		return status.Error(codes.InvalidArgument, "Different AppID")
	}

	return nil
}

func Validate(info *npool.AccountReq) error {
	return validate(info)
}
