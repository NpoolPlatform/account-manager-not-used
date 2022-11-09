//nolint:nolintlint,dupl
package platform

import (
	"context"

	converter "github.com/NpoolPlatform/account-manager/pkg/converter/platform"
	crud "github.com/NpoolPlatform/account-manager/pkg/crud/platform"
	commontracer "github.com/NpoolPlatform/account-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/account-manager/pkg/tracer/platform"

	constant "github.com/NpoolPlatform/account-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/platform"

	"github.com/google/uuid"
)

func (s *Server) CreateAccount(ctx context.Context, in *npool.CreateAccountRequest) (*npool.CreateAccountResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAccount")
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
		return &npool.CreateAccountResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "platform", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create platform: %v", err.Error())
		return &npool.CreateAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAccountResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateAccounts(ctx context.Context, in *npool.CreateAccountsRequest) (*npool.CreateAccountsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAccounts")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateAccountsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	err = duplicate(in.GetInfos())
	if err != nil {
		return &npool.CreateAccountsResponse{}, err
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "platform", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create platforms: %v", err)
		return &npool.CreateAccountsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAccountsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateAccount(ctx context.Context, in *npool.UpdateAccountRequest) (*npool.UpdateAccountResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAccount")
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
		return &npool.UpdateAccountResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "platform", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create platform: %v", err.Error())
		return &npool.UpdateAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAccountResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAccount(ctx context.Context, in *npool.GetAccountRequest) (*npool.GetAccountResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAccount")
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
		return &npool.GetAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "platform", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get platform: %v", err)
		return &npool.GetAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAccountResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAccountOnly(ctx context.Context, in *npool.GetAccountOnlyRequest) (*npool.GetAccountOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAccountOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "platform", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get platforms: %v", err)
		return &npool.GetAccountOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAccountOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAccounts(ctx context.Context, in *npool.GetAccountsRequest) (*npool.GetAccountsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAccounts")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "platform", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get platforms: %v", err)
		return &npool.GetAccountsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAccountsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistAccount(ctx context.Context, in *npool.ExistAccountRequest) (*npool.ExistAccountResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAccount")
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
		return &npool.ExistAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "platform", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check platform: %v", err)
		return &npool.ExistAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAccountResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistAccountConds(ctx context.Context,
	in *npool.ExistAccountCondsRequest) (*npool.ExistAccountCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAccountConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "platform", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check platform: %v", err)
		return &npool.ExistAccountCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAccountCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountAccounts(ctx context.Context, in *npool.CountAccountsRequest) (*npool.CountAccountsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountAccounts")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "platform", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count platforms: %v", err)
		return &npool.CountAccountsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAccountsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteAccount(ctx context.Context, in *npool.DeleteAccountRequest) (*npool.DeleteAccountResponse, error) {
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
		return &npool.DeleteAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "goodbenefit", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete goodbenefit: %v", err)
		return &npool.DeleteAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAccountResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
