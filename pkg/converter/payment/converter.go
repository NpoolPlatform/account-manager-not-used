package payment

import (
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/payment"
)

func Ent2Grpc(row *ent.Payment) *npool.Account {
	if row == nil {
		return nil
	}

	return &npool.Account{
		ID:            row.ID.String(),
		CoinTypeID:    row.CoinTypeID.String(),
		AccountID:     row.AccountID.String(),
		Idle:          row.Idle,
		OccupiedBy:    npool.OccupiedBy(npool.OccupiedBy_value[row.OccupiedBy]),
		CollectingTID: row.CollectingTid.String(),
		CreatedAt:     row.CreatedAt,
		AvailableAt:   row.AvailableAt,
		UpdatedAt:     row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Payment) []*npool.Account {
	infos := []*npool.Account{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
