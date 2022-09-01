package transfer

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	val "github.com/NpoolPlatform/message/npool"

	testinit "github.com/NpoolPlatform/account-manager/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/transfer"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var kycData = npool.Transfer{
	ID:           uuid.NewString(),
	AppID:        uuid.NewString(),
	UserID:       uuid.NewString(),
	TargetUserID: uuid.NewString(),
}

var (
	kycInfo = npool.TransferReq{
		ID:           &kycData.ID,
		AppID:        &kycData.AppID,
		UserID:       &kycData.UserID,
		TargetUserID: &kycData.TargetUserID,
	}
)

var info *npool.Transfer

func createTransfer(t *testing.T) {
	var err error
	info, err = CreateTransfer(context.Background(), &kycInfo)
	if assert.Nil(t, err) {
		kycData.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &kycData)
	}
}

func createTransfers(t *testing.T) {
	kycDatas := []npool.Transfer{
		{
			ID:           uuid.NewString(),
			AppID:        uuid.NewString(),
			UserID:       uuid.NewString(),
			TargetUserID: uuid.NewString(),
		},
		{
			ID:           uuid.NewString(),
			AppID:        uuid.NewString(),
			UserID:       uuid.NewString(),
			TargetUserID: uuid.NewString(),
		},
	}

	Transfers := []*npool.TransferReq{}
	for key := range kycDatas {
		Transfers = append(Transfers, &npool.TransferReq{
			ID:           &kycDatas[key].ID,
			AppID:        &kycDatas[key].AppID,
			UserID:       &kycDatas[key].UserID,
			TargetUserID: &kycDatas[key].TargetUserID,
		})
	}

	infos, err := CreateTransfers(context.Background(), Transfers)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func getTransfer(t *testing.T) {
	var err error
	info, err = GetTransfer(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &kycData)
	}
}

func getTransfers(t *testing.T) {
	infos, total, err := GetTransfers(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &kycData)
	}
}

func getTransferOnly(t *testing.T) {
	var err error
	info, err = GetTransferOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &kycData)
	}
}

func countTransfers(t *testing.T) {
	count, err := CountTransfers(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, uint32(1))
	}
}

func existTransfer(t *testing.T) {
	exist, err := ExistTransfer(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existTransferConds(t *testing.T) {
	exist, err := ExistTransferConds(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteTransfer(t *testing.T) {
	info, err := DeleteTransfer(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &kycData)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createTransfer", createTransfer)
	t.Run("createTransfers", createTransfers)
	t.Run("getTransfer", getTransfer)
	t.Run("getTransfers", getTransfers)
	t.Run("getTransferOnly", getTransferOnly)
	t.Run("existTransfer", existTransfer)
	t.Run("existTransferConds", existTransferConds)
	t.Run("count", countTransfers)
	t.Run("delete", deleteTransfer)
}
