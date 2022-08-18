package deposit

import (
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/deposit"
)

func Ent2Grpc(row *ent.Deposit) *npool.Account {
	if row == nil {
		return nil
	}

	return &npool.Account{
		ID:            row.ID.String(),
		AppID:         row.AppID.String(),
		UserID:        row.UserID.String(),
		CoinTypeID:    row.CoinTypeID.String(),
		AccountID:     row.AccountID.String(),
		Incoming:      row.Incoming.String(),
		Outcoming:     row.Outcoming.String(),
		CollectingTID: row.CollectingTid.String(),
		ScannableAt:   row.ScannableAt,
		CreatedAt:     row.CreatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Deposit) []*npool.Account {
	infos := []*npool.Account{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
