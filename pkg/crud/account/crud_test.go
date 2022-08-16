package account

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	testinit "github.com/NpoolPlatform/account-manager/pkg/testinit"
	valuedef "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
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

var entity = ent.Account{
	ID:         uuid.New(),
	CoinTypeID: uuid.New(),
	UsedFor:    npool.AccountUsedFor_GoodPayment.String(),
	Address:    uuid.New().String(),
}

var (
	id         = entity.ID.String()
	coinTypeID = entity.CoinTypeID.String()
	usedFor    = npool.AccountUsedFor_GoodPayment
	address    = entity.Address

	req = npool.AccountReq{
		ID:         &id,
		CoinTypeID: &coinTypeID,
		UsedFor:    &usedFor,
		Address:    &address,
	}
)

var info *ent.Account

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		entity.UpdatedAt = info.UpdatedAt
		entity.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), entity.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.Account{
		{
			ID:         uuid.New(),
			CoinTypeID: uuid.New(),
			UsedFor:    npool.AccountUsedFor_GoodPayment.String(),
			Address:    uuid.New().String(),
		},
		{
			ID:         uuid.New(),
			CoinTypeID: uuid.New(),
			UsedFor:    npool.AccountUsedFor_GoodPayment.String(),
			Address:    uuid.New().String(),
		},
	}

	reqs := []*npool.AccountReq{}
	for _, _entity := range entities {
		_id := _entity.ID.String()
		_coinTypeID := _entity.CoinTypeID.String()
		_usedFor := npool.AccountUsedFor_GoodPayment
		_address := _entity.Address

		reqs = append(reqs, &npool.AccountReq{
			ID:         &_id,
			CoinTypeID: &_coinTypeID,
			UsedFor:    &_usedFor,
			Address:    &_address,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func update(t *testing.T) {
	active := true

	req.Active = &active
	req.Locked = &active
	req.Blocked = &active

	entity.Active = active
	entity.Locked = active
	entity.Blocked = active

	info, err := Update(context.Background(), &req)
	if assert.Nil(t, err) {
		entity.UpdatedAt = info.UpdatedAt
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
