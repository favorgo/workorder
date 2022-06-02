// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.2

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type OrderHTTPServer interface {
	Purchase(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterOrderHTTPServer(s *http.Server, srv OrderHTTPServer) {
	r := s.Route("/")
	r.GET("/workorder/{name}", _Order_Purchase0_HTTP_Handler(srv))
}

func _Order_Purchase0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/order.v1.Order/Purchase")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Purchase(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

type OrderHTTPClient interface {
	Purchase(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
}

type OrderHTTPClientImpl struct {
	cc *http.Client
}

func NewOrderHTTPClient(client *http.Client) OrderHTTPClient {
	return &OrderHTTPClientImpl{client}
}

func (c *OrderHTTPClientImpl) Purchase(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/workorder/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/order.v1.Order/Purchase"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}