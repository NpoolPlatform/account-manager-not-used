package account

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
)

func trace(span trace1.Span, in *npool.AccountReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("CoinTypeID.%v", index), in.GetCoinTypeID()),
		attribute.String(fmt.Sprintf("Address.%v", index), in.GetAddress()),
		attribute.String(fmt.Sprintf("UsedFor.%v", index), in.GetUsedFor().String()),
		attribute.Bool(fmt.Sprintf("PlatformHoldPrivateKey.%v", index), in.GetPlatformHoldPrivateKey()),
		attribute.Bool(fmt.Sprintf("Active.%v", index), in.GetActive()),
		attribute.Bool(fmt.Sprintf("Locked.%v", index), in.GetLocked()),
		attribute.Bool(fmt.Sprintf("Blocked.%v", index), in.GetBlocked()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.AccountReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("CoinTypeID.Op", in.GetCoinTypeID().GetOp()),
		attribute.String("CoinTypeID.Value", in.GetCoinTypeID().GetValue()),
		attribute.String("Address.Op", in.GetAddress().GetOp()),
		attribute.String("Address.Value", in.GetAddress().GetValue()),
		attribute.String("UsedFor.Op", in.GetUsedFor().GetOp()),
		attribute.String("UsedFor.Value", npool.AccountUsedFor(in.GetUsedFor().GetValue()).String()),
		attribute.String("PlatformHoldPrivateKey.Op", in.GetPlatformHoldPrivateKey().GetOp()),
		attribute.Bool("PlatformHoldPrivateKey.Value", in.GetPlatformHoldPrivateKey().GetValue()),
		attribute.String("Active.Op", in.GetActive().GetOp()),
		attribute.Bool("Active.Value", in.GetActive().GetValue()),
		attribute.String("Locked.Op", in.GetLocked().GetOp()),
		attribute.Bool("Locked.Value", in.GetLocked().GetValue()),
		attribute.String("Blocked.Op", in.GetBlocked().GetOp()),
		attribute.Bool("Blocked.Value", in.GetBlocked().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.AccountReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
