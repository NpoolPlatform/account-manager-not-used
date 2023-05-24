package transfer

import (
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/transfer"
)

func Ent2Grpc(row *ent.Transfer) *npool.Transfer {
	if row == nil {
		return nil
	}

	return &npool.Transfer{
		ID:           row.ID.String(),
		AppID:        row.AppID.String(),
		UserID:       row.UserID.String(),
		TargetUserID: row.TargetUserID.String(),
		CreatedAt:    row.CreatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Transfer) []*npool.Transfer {
	infos := []*npool.Transfer{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
