package endpoint

import (
	"github.com/king19800105/go-kit-learn/demo1/pkg/service"
	"github.com/go-kit/kit/endpoint"
	"context"
)

// 请求参数格式
type CreateRequest struct {
	OrderId string `json:"orderId"`
}

// 响应参数格式
type CreateResponse struct {
	Code int   `json:"code"`
	Err  error `json:"err"`
}

// create操作端点
func makeCreateEndpoint(svc service.OrderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// 模拟设置ctx对象
		ctx = context.WithValue(ctx, "ID", "123456")
		req := request.(CreateRequest)
		code, err := svc.Create(ctx, req.OrderId)

		return CreateResponse{
			Code: code,
			Err:  err,
		}, nil
	}
}
