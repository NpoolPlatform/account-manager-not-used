package platform

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/platform"
)

func trace(span trace1.Span, in *npool.AccountReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AccountID.%v", index), in.GetAccountID()),
		attribute.String(fmt.Sprintf("UsedFor.%v", index), in.GetUsedFor().String()),
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
		attribute.String("AccountID.Op", in.GetAccountID().GetOp()),
		attribute.String("AccountID.Value", in.GetAccountID().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.AccountReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
