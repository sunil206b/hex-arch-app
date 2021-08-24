package grpc

import (
	"context"
	"github.com/sunil206b/hex-arc-app/internal/adapters/framework/left/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ad *Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}
	answer, err := ad.api.GetAddition(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{
		Value: answer,
	}
	return ans, nil
}

func (ad *Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}
	answer, err := ad.api.GetSubtraction(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{
		Value: answer,
	}
	return ans, nil
}

func (ad *Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}
	answer, err := ad.api.GetAddition(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{
		Value: answer,
	}
	return ans, nil
}

func (ad *Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}
	answer, err := ad.api.GetAddition(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{
		Value: answer,
	}
	return ans, nil
}
