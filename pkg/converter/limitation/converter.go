package limitation

import (
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/limitation"
)

func Ent2Grpc(row *ent.Limitation) *npool.Limitation {
	if row == nil {
		return nil
	}

	return &npool.Limitation{
		ID:         row.ID.String(),
		CoinTypeID: row.CoinTypeID.String(),
		Limitation: npool.LimitationType(npool.LimitationType_value[row.Limitation]),
		Amount:     row.Amount.String(),
		CreatedAt:  row.CreatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Limitation) []*npool.Limitation {
	infos := []*npool.Limitation{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
