package service

import (
	"context"
	"time"

	v1 "workorder/api/helloworld/v1"
	"workorder/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedOrderServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) Purchase(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	time.Sleep(200 * time.Millisecond)
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
