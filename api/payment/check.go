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

	return nil
}

func duplicate(infos []*npool.AccountReq) error {
	keys := map[string]struct{}{}

	for _, info := range infos {
		if err := validate(info); err != nil {
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Infos has invalid element %v", err))
		}

		key := fmt.Sprintf("%v", info.GetAccountID())
		if _, ok := keys[key]; ok {
			return status.Error(codes.InvalidArgument, "Infos has duplicate AccountID")
		}

		keys[key] = struct{}{}
	}

	return nil
}
