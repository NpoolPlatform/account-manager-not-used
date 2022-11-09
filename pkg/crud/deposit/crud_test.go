package deposit

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	testinit "github.com/NpoolPlatform/account-manager/pkg/testinit"
	valuedef "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/deposit"

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

var entity = ent.Deposit{
	ID:         uuid.New(),
	AppID:      uuid.New(),
	UserID:     uuid.New(),
	CoinTypeID: uuid.New(),
	AccountID:  uuid.New(),
}

var (
	id         = entity.ID.String()
	appID      = entity.AppID.String()
	depositID  = entity.UserID.String()
	coinTypeID = entity.CoinTypeID.String()
	accountID  = entity.AccountID.String()

	req = npool.AccountReq{
		ID:         &id,
		AppID:      &appID,
		UserID:     &depositID,
		CoinTypeID: &coinTypeID,
		AccountID:  &accountID,
	}
)

var info *ent.Deposit

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		entity.UpdatedAt = info.UpdatedAt
		entity.ScannableAt = info.ScannableAt
		entity.CreatedAt = info.CreatedAt
		entity.ScannableAt = info.ScannableAt
		assert.Equal(t, info.String(), entity.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.Deposit{
		{
			ID:         uuid.New(),
			AppID:      uuid.New(),
			UserID:     uuid.New(),
			CoinTypeID: uuid.New(),
			AccountID:  uuid.New(),
		},
		{
			ID:         uuid.New(),
			AppID:      uuid.New(),
			UserID:     uuid.New(),
			CoinTypeID: uuid.New(),
			AccountID:  uuid.New(),
		},
	}

	reqs := []*npool.AccountReq{}
	for _, _entity := range entities {
		_id := _entity.ID.String()
		_appID := _entity.AppID.String()
		_depositID := _entity.UserID.String()
		_coinTypeID := _entity.CoinTypeID.String()
		_accountID := _entity.AccountID.String()

		reqs = append(reqs, &npool.AccountReq{
			ID:         &_id,
			AppID:      &_appID,
			UserID:     &_depositID,
			CoinTypeID: &_coinTypeID,
			AccountID:  &_accountID,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func update(t *testing.T) {
	collectingTID := uuid.New().String()
	scannableAt := uint32(time.Now().Unix() + 10000)

	req.CollectingTID = &collectingTID
	req.ScannableAt = &scannableAt

	entity.CollectingTid = uuid.MustParse(collectingTID)
	entity.ScannableAt = scannableAt

	info, err := Update(context.Background(), &req)
	if assert.Nil(t, err) {
		entity.UpdatedAt = info.UpdatedAt
		entity.ScannableAt = info.ScannableAt
		assert.Equal(t, info.String(), entity.String())
	}

	incoming := "9999.8988"
	req.Incoming = &incoming
	entity.Incoming = decimal.RequireFromString(incoming)

	info, err = AddFields(context.Background(), &req)
	if assert.Nil(t, err) {
		entity.UpdatedAt = info.UpdatedAt
		entity.ScannableAt = info.ScannableAt
		assert.Equal(t, info.String(), entity.String())
	}

	req.Incoming = &incoming
	entity.Incoming = decimal.RequireFromString(incoming).Add(entity.Incoming)
	info, err = AddFields(context.Background(), &req)
	if assert.Nil(t, err) {
		entity.UpdatedAt = info.UpdatedAt
		entity.ScannableAt = info.ScannableAt
		assert.Equal(t, info.String(), entity.String())
	}

	outcoming := "9999.999"
	req.Outcoming = &outcoming
	req.Incoming = nil
	entity.Outcoming = decimal.RequireFromString(outcoming)
	info, err = AddFields(context.Background(), &req)
	if assert.Nil(t, err) {
		entity.UpdatedAt = info.UpdatedAt
		entity.ScannableAt = info.ScannableAt
		assert.Equal(t, info.String(), entity.String())
	}

	outcoming = "999.999"
	req.Outcoming = &outcoming
	entity.Outcoming = decimal.RequireFromString(outcoming).Add(entity.Outcoming)
	info, err = AddFields(context.Background(), &req)
	if assert.Nil(t, err) {
		entity.UpdatedAt = info.UpdatedAt
		entity.ScannableAt = info.ScannableAt
		assert.Equal(t, info.String(), entity.String())
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), entity.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), entity.String())
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		if assert.Equal(t, total, 1) {
			assert.Equal(t, infos[0].String(), entity.String())
		}
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), entity.String())
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, uint32(1))
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), entity.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteA(t *testing.T) {
	info, err := Delete(context.Background(), entity.ID)
	if assert.Nil(t, err) {
		entity.DeletedAt = info.DeletedAt
		entity.UpdatedAt = info.UpdatedAt
		entity.ScannableAt = info.ScannableAt
		assert.Equal(t, info.String(), entity.String())
	}
}

func TestAccount(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("update", update)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("count", count)
	t.Run("delete", deleteA)
}
