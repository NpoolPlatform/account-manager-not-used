package limitation

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
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/limitation"
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

var accountData = npool.Limitation{
	ID:         uuid.NewString(),
	CoinTypeID: uuid.NewString(),
	Limitation: npool.LimitationType_GasFeedingAmount,
	Amount:     "10",
}

var (
	accountInfo = npool.LimitationReq{
		ID:         &accountData.ID,
		CoinTypeID: &accountData.CoinTypeID,
		Limitation: &accountData.Limitation,
		Amount:     &accountData.Amount,
	}
)

var info *npool.Limitation

func createLimitation(t *testing.T) {
	var err error
	info, err = CreateLimitation(context.Background(), &accountInfo)
	if assert.Nil(t, err) {
		accountData.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &accountData)
	}
}

func updateLimitation(t *testing.T) {
	var err error
	info, err = UpdateLimitation(context.Background(), &accountInfo)
	if assert.Nil(t, err) {
		accountData.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &accountData)
	}
}

func createLimitations(t *testing.T) {
	accountDatas := []npool.Limitation{
		{
			ID:         uuid.NewString(),
			CoinTypeID: uuid.NewString(),
			Limitation: npool.LimitationType_GasFeedingAmount,
			Amount:     "10",
		},
		{
			ID:         uuid.NewString(),
			CoinTypeID: uuid.NewString(),
			Limitation: npool.LimitationType_GasFeedingAmount,
			Amount:     "10",
		},
	}

	Limitations := []*npool.LimitationReq{}
	for key := range accountDatas {
		Limitations = append(Limitations, &npool.LimitationReq{
			ID:         &accountDatas[key].ID,
			CoinTypeID: &accountDatas[key].CoinTypeID,
			Limitation: &accountDatas[key].Limitation,
			Amount:     &accountDatas[key].Amount,
		})
	}

	infos, err := CreateLimitations(context.Background(), Limitations)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func getLimitation(t *testing.T) {
	var err error
	info, err = GetLimitation(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &accountData)
	}
}

func getLimitations(t *testing.T) {
	infos, total, err := GetLimitations(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &accountData)
	}
}

func getLimitationOnly(t *testing.T) {
	var err error
	info, err = GetLimitationOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &accountData)
	}
}

func countLimitations(t *testing.T) {
	count, err := CountLimitations(context.Background(),
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

func existLimitation(t *testing.T) {
	exist, err := ExistLimitation(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existLimitationConds(t *testing.T) {
	exist, err := ExistLimitationConds(context.Background(),
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

func deleteLimitation(t *testing.T) {
	info, err := DeleteLimitation(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &accountData)
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

	t.Run("createLimitation", createLimitation)
	t.Run("updateLimitation", updateLimitation)
	t.Run("createLimitations", createLimitations)
	t.Run("getLimitation", getLimitation)
	t.Run("getLimitations", getLimitations)
	t.Run("getLimitationOnly", getLimitationOnly)
	t.Run("existLimitation", existLimitation)
	t.Run("existLimitationConds", existLimitationConds)
	t.Run("count", countLimitations)
	t.Run("delete", deleteLimitation)
}
