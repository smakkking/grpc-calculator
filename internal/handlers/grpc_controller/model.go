package grpc_controller

import (
	"context"
	"fmt"
	"grpc_calc/pkg/sdk/go/calcv1"

	"google.golang.org/grpc"
)

type serverAPI struct {
	// вот эта хрень нужна, если вы неполностью реализовали service,
	// там просто заглушки вместо методов стоят
	calcv1.UnimplementedCalculatorServer
}

// вот эту хрень нужно будет вызывать в main, чтобы получить grpc сервер работающий
func Register(gRPC *grpc.Server) {
	calcv1.RegisterCalculatorServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Add(ctx context.Context, req *calcv1.Request) (*calcv1.Responce, error) {
	res := req.A + req.B
	return &calcv1.Responce{Result: res}, nil
}

func (s *serverAPI) Sub(ctx context.Context, req *calcv1.Request) (*calcv1.Responce, error) {
	res := req.A - req.B
	return &calcv1.Responce{Result: res}, nil
}

func (s *serverAPI) Mul(ctx context.Context, req *calcv1.Request) (*calcv1.Responce, error) {
	res := req.A * req.B
	return &calcv1.Responce{Result: res}, nil
}

func (s *serverAPI) Div(ctx context.Context, req *calcv1.Request) (*calcv1.Responce, error) {
	if req.B == 0 {
		return nil, fmt.Errorf("division by 0")
	}

	res := req.A / req.B
	return &calcv1.Responce{Result: res}, nil
}
