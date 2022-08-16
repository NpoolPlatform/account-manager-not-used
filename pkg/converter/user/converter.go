package user

import (
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	account "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/user"
)

func Ent2Grpc(row *ent.User) *npool.Account {
	if row == nil {
		return nil
	}

	return &npool.Account{
		ID:         row.ID.String(),
		AppID:      row.AppID.String(),
		UserID:     row.UserID.String(),
		CoinTypeID: row.CoinTypeID.String(),
		AccountID:  row.AccountID.String(),
		UsedFor:    account.AccountUsedFor(account.AccountUsedFor_value[row.UsedFor]),
		Labels:     row.Labels,
		CreatedAt:  row.CreatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.User) []*npool.Account {
	infos := []*npool.Account{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
