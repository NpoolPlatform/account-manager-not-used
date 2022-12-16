package account

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/account-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/account-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/account-manager/pkg/tracer/account"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/account-manager/pkg/db"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/account"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"

	"github.com/google/uuid"
)

func CreateSet(c *ent.AccountCreate, in *npool.AccountReq) *ent.AccountCreate {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.CoinTypeID != nil {
		c.SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID()))
	}
	if in.Address != nil {
		c.SetAddress(in.GetAddress())
	}
	if in.UsedFor != nil {
		c.SetUsedFor(in.GetUsedFor().String())
	}
	if in.PlatformHoldPrivateKey != nil {
		c.SetPlatformHoldPrivateKey(in.GetPlatformHoldPrivateKey())
	}
	if in.Active != nil {
		c.SetActive(in.GetActive())
	}
	if in.Locked != nil {
		c.SetLocked(in.GetLocked())
	}
	if in.LockedBy != nil {
		c.SetLockedBy(in.GetLockedBy().String())
	}
	if in.Blocked != nil {
		c.SetBlocked(in.GetBlocked())
	}
	return c
}

func Create(ctx context.Context, in *npool.AccountReq) (*ent.Account, error) {
	var info *ent.Account
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.Account.Create()
		info, err = CreateSet(c, in).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.AccountReq) ([]*ent.Account, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceMany(span, in)

	rows := []*ent.Account{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.AccountCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.Account.Create(), info)
		}
		rows, err = tx.Account.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func UpdateSet(info *ent.Account, in *npool.AccountReq) *ent.AccountUpdateOne {
	u := info.Update()

	if in.Active != nil {
		u.SetActive(in.GetActive())
	}
	if in.Locked != nil {
		u.SetLocked(in.GetLocked())
	}
	if in.LockedBy != nil {
		u.SetLockedBy(in.GetLockedBy().String())
	}
	if in.Blocked != nil {
		u.SetBlocked(in.GetBlocked())
	}
	if in.DeletedAt != nil {
		u.SetDeletedAt(in.GetDeletedAt())
	}

	return u
}

func Update(ctx context.Context, in *npool.AccountReq) (*ent.Account, error) {
	var info *ent.Account
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Update")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err = tx.Account.Query().Where(account.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return err
		}

		info, err = UpdateSet(info, in).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.Account, error) {
	var info *ent.Account
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Row")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Account.Query().Where(account.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AccountQuery, error) { // nolint
	stm := cli.Account.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(account.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	if conds.CoinTypeID != nil {
		switch conds.GetCoinTypeID().GetOp() {
		case cruder.EQ:
			stm.Where(account.CoinTypeID(uuid.MustParse(conds.GetCoinTypeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	if conds.Address != nil {
		switch conds.GetAddress().GetOp() {
		case cruder.EQ:
			stm.Where(account.Address(conds.GetAddress().GetValue()))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	if conds.UsedFor != nil {
		switch conds.GetUsedFor().GetOp() {
		case cruder.EQ:
			stm.Where(account.UsedFor(npool.AccountUsedFor(conds.GetUsedFor().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	if conds.PlatformHoldPrivateKey != nil {
		switch conds.GetPlatformHoldPrivateKey().GetOp() {
		case cruder.EQ:
			stm.Where(account.PlatformHoldPrivateKey(conds.GetPlatformHoldPrivateKey().GetValue()))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	if conds.Active != nil {
		switch conds.GetActive().GetOp() {
		case cruder.EQ:
			stm.Where(account.Active(conds.GetActive().GetValue()))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	if conds.Locked != nil {
		switch conds.GetLocked().GetOp() {
		case cruder.EQ:
			stm.Where(account.Locked(conds.GetLocked().GetValue()))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	if conds.LockedBy != nil {
		switch conds.GetLockedBy().GetOp() {
		case cruder.EQ:
			stm.Where(account.LockedBy(npool.LockedBy(conds.GetLockedBy().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	if conds.Blocked != nil {
		switch conds.GetBlocked().GetOp() {
		case cruder.EQ:
			stm.Where(account.Blocked(conds.GetBlocked().GetValue()))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	if conds.IDs != nil {
		switch conds.GetIDs().GetOp() {
		case cruder.IN:
			ids := []uuid.UUID{}
			for _, id := range conds.GetIDs().GetValue() {
				ids = append(ids, uuid.MustParse(id))
			}
			stm.Where(account.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Account, int, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)
	span = commontracer.TraceOffsetLimit(span, offset, limit)

	rows := []*ent.Account{}
	var total int
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(account.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.Account, error) {
	var info *ent.Account
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Count(ctx context.Context, conds *npool.Conds) (uint32, error) {
	var err error
	var total int

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Count")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return uint32(total), nil
}

func Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Exist")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.Account.Query().Where(account.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func Delete(ctx context.Context, id uuid.UUID) (*ent.Account, error) {
	var info *ent.Account
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Delete")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Account.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
