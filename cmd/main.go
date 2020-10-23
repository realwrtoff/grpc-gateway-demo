package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	echo "github.com/realwrtoff/rest_grpc/proto/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type EchoService struct{}

func (s *EchoService) Echo(ctx context.Context, req *echo.EchoReq) (*echo.EchoRes, error) {
	return &echo.EchoRes{Name: req.Name, Age: req.Age + 1}, nil
}

type CalService struct{}

func (s *CalService) Cal(ctx context.Context, req *echo.CalReq) (*echo.CalRes, error) {
	var result int64
	switch req.Info.Op {
	case "+":
		result = req.Info.A + req.Info.B
	case "-":
		result = req.Info.A - req.Info.B
	default:
		return nil, status.Errorf(codes.InvalidArgument, "op should in ['+', '-']")
	}

	return &echo.CalRes{
		Result: result,
		Uid:    req.Uid,
	}, nil
}


func main() {
	mux := runtime.NewServeMux()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	if err := echo.RegisterEchoServiceHandlerServer(ctx, mux, &EchoService{}); err != nil {
		panic(err)
	}
	if err := echo.RegisterCalServiceHandlerServer(ctx, mux, &CalService{}); err != nil {
		panic(err)
	}
	if err := http.ListenAndServe(":21680", mux); err != nil {
		panic(err)
	}
}
