package goodbenefit

import (
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/goodbenefit"
)

func Ent2Grpc(row *ent.GoodBenefit) *npool.Account {
	if row == nil {
		return nil
	}

	return &npool.Account{
		ID:            row.ID.String(),
		GoodID:        row.GoodID.String(),
		AccountID:     row.AccountID.String(),
		Backup:        row.Backup,
		TransactionID: row.TransactionID.String(),
		IntervalHours: row.IntervalHours,
	}
}

func Ent2GrpcMany(rows []*ent.GoodBenefit) []*npool.Account {
	infos := []*npool.Account{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
