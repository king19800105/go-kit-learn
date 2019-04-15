package endpoint

import (
	"github.com/king19800105/go-kit-learn/demo2/pkg/service"
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/king19800105/go-kit-learn/demo2/pkg/entity"
)

// 请求参数格式
type CreateRequest struct {
	OrderId string `json:"orderId"`
}

// 响应参数格式
type CreateResponse struct {
	entity.Order
	err error
}

// create操作端点
func makeCreateEndpoint(svc service.OrderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		order, err := svc.Create(ctx, req.OrderId)

		return CreateResponse{
			order,
			err,
		}, nil
	}
}

// 错误获取
func (rs CreateResponse) Failed() error {
	return rs.err
}
