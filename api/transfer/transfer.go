//nolint:nolintlint,dupl
package transfer

import (
	"context"

	converter "github.com/NpoolPlatform/account-manager/pkg/converter/transfer"
	crud "github.com/NpoolPlatform/account-manager/pkg/crud/transfer"
	commontracer "github.com/NpoolPlatform/account-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/account-manager/pkg/tracer/transfer"

	constant "github.com/NpoolPlatform/account-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/transfer"

	"github.com/google/uuid"
)

func (s *Server) CreateTransfer(ctx context.Context, in *npool.CreateTransferRequest) (*npool.CreateTransferResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateTransfer")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	err = validate(in.GetInfo())
	if err != nil {
		return &npool.CreateTransferResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "transfer", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create transfer: %v", err.Error())
		return &npool.CreateTransferResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTransferResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateTransfers(ctx context.Context, in *npool.CreateTransfersRequest) (*npool.CreateTransfersResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateTransfers")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateTransfersResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "transfer", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create transfers: %v", err)
		return &npool.CreateTransfersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTransfersResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) GetTransfer(ctx context.Context, in *npool.GetTransferRequest) (*npool.GetTransferResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTransfer")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetTransferResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "transfer", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get transfer: %v", err)
		return &npool.GetTransferResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTransferResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetTransferOnly(ctx context.Context, in *npool.GetTransferOnlyRequest) (*npool.GetTransferOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTransferOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "transfer", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get transfers: %v", err)
		return &npool.GetTransferOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTransferOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetTransfers(ctx context.Context, in *npool.GetTransfersRequest) (*npool.GetTransfersResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTransfers")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "transfer", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get transfers: %v", err)
		return &npool.GetTransfersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTransfersResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistTransfer(ctx context.Context, in *npool.ExistTransferRequest) (*npool.ExistTransferResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistTransfer")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistTransferResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "transfer", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check transfer: %v", err)
		return &npool.ExistTransferResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTransferResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistTransferConds(ctx context.Context,
	in *npool.ExistTransferCondsRequest) (*npool.ExistTransferCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistTransferConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "transfer", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check transfer: %v", err)
		return &npool.ExistTransferCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTransferCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountTransfers(ctx context.Context, in *npool.CountTransfersRequest) (*npool.CountTransfersResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountTransfers")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "transfer", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count transfers: %v", err)
		return &npool.CountTransfersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountTransfersResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteTransfer(ctx context.Context, in *npool.DeleteTransferRequest) (*npool.DeleteTransferResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteGoodBenefit")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.DeleteTransferResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "goodbenefit", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete goodbenefit: %v", err)
		return &npool.DeleteTransferResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteTransferResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
