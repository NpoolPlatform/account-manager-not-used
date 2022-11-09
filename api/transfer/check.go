package transfer

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/transfer"

	"github.com/google/uuid"
)

func validate(info *npool.TransferReq) error {
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

	if info.TargetUserID == nil {
		logger.Sugar().Errorw("validate", "TargetUserID", info.TargetUserID)
		return status.Error(codes.InvalidArgument, "TargetUserID is empty")
	}

	if _, err := uuid.Parse(info.GetTargetUserID()); err != nil {
		logger.Sugar().Errorw("validate", "TargetUserID", info.GetTargetUserID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("TargetUserID is invalid: %v", err))
	}

	return nil
}

func Validate(info *npool.TransferReq) error {
	return validate(info)
}
