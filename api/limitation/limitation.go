//nolint:nolintlint,dupl
package limitation

import (
	"context"

	converter "github.com/NpoolPlatform/account-manager/pkg/converter/limitation"
	crud "github.com/NpoolPlatform/account-manager/pkg/crud/limitation"
	commontracer "github.com/NpoolPlatform/account-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/account-manager/pkg/tracer/limitation"

	constant "github.com/NpoolPlatform/account-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/limitation"

	"github.com/google/uuid"
)

func (s *Server) CreateLimitation(ctx context.Context, in *npool.CreateLimitationRequest) (*npool.CreateLimitationResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateLimitation")
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
		return &npool.CreateLimitationResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "limitation", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create limitation: %v", err.Error())
		return &npool.CreateLimitationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateLimitationResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateLimitations(ctx context.Context, in *npool.CreateLimitationsRequest) (*npool.CreateLimitationsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateLimitations")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateLimitationsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	err = duplicate(in.GetInfos())
	if err != nil {
		return &npool.CreateLimitationsResponse{}, err
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "limitation", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create limitations: %v", err)
		return &npool.CreateLimitationsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateLimitationsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateLimitation(ctx context.Context, in *npool.UpdateLimitationRequest) (*npool.UpdateLimitationResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateLimitation")
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
		return &npool.UpdateLimitationResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "limitation", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create limitation: %v", err.Error())
		return &npool.UpdateLimitationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateLimitationResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetLimitation(ctx context.Context, in *npool.GetLimitationRequest) (*npool.GetLimitationResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetLimitation")
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
		return &npool.GetLimitationResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "limitation", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get limitation: %v", err)
		return &npool.GetLimitationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetLimitationResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetLimitationOnly(ctx context.Context, in *npool.GetLimitationOnlyRequest) (*npool.GetLimitationOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetLimitationOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "limitation", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get limitations: %v", err)
		return &npool.GetLimitationOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetLimitationOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetLimitations(ctx context.Context, in *npool.GetLimitationsRequest) (*npool.GetLimitationsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetLimitations")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "limitation", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get limitations: %v", err)
		return &npool.GetLimitationsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetLimitationsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistLimitation(ctx context.Context, in *npool.ExistLimitationRequest) (*npool.ExistLimitationResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistLimitation")
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
		return &npool.ExistLimitationResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "limitation", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check limitation: %v", err)
		return &npool.ExistLimitationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistLimitationResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistLimitationConds(ctx context.Context,
	in *npool.ExistLimitationCondsRequest) (*npool.ExistLimitationCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistLimitationConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "limitation", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check limitation: %v", err)
		return &npool.ExistLimitationCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistLimitationCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountLimitations(ctx context.Context, in *npool.CountLimitationsRequest) (*npool.CountLimitationsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountLimitations")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "limitation", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count limitations: %v", err)
		return &npool.CountLimitationsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountLimitationsResponse{
		Info: total,
	}, nil
}
