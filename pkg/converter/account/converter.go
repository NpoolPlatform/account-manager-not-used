package account

import (
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
)

func Ent2Grpc(row *ent.Account) *npool.Account {
	if row == nil {
		return nil
	}

	return &npool.Account{
		ID:                     row.ID.String(),
		CoinTypeID:             row.CoinTypeID.String(),
		Address:                row.Address,
		UsedFor:                npool.AccountUsedFor(npool.AccountUsedFor_value[row.UsedFor]),
		PlatformHoldPrivateKey: row.PlatformHoldPrivateKey,
		Active:                 row.Active,
		Locked:                 row.Locked,
		LockedBy:               npool.LockedBy(npool.LockedBy_value[row.LockedBy]),
		Blocked:                row.Blocked,
	}
}

func Ent2GrpcMany(rows []*ent.Account) []*npool.Account {
	infos := []*npool.Account{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
