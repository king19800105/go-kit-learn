package service

import (
	"context"
	"github.com/king19800105/go-kit-learn/demo1/pkg/entity"
	"github.com/king19800105/go-kit-learn/demo1/msg"
)

// 服务抽象
type OrderService interface {
	Create(ctx context.Context, orderId string) (entity.Order, error)
}

// 订单结构
type baseOrderService struct{}

// todo... 再加一个delete状态，只返回code，还有使用nop的空结构
// todo... 添加一个get，返回[]io.Order
// todo... 添加一个
// 创建订单
func (os baseOrderService) Create(ctx context.Context, orderId string) (o entity.Order, err error) {
	if "" == orderId {
		return o, msg.New(msg.ORDER_NO_EMPTY)
	}

	o = entity.Order{
		Id:     "#" + orderId,
		Source: "APP",
		IsPay:  1,
	}

	return o, nil
}

// 服务对象实例化，并且组装中间件
func New(middleware []Middleware) OrderService {
	var svc = getBaseService()

	for _, m := range middleware {
		svc = m(svc)
	}

	return svc
}

// 获取当前实例
func getBaseService() OrderService {
	return &baseOrderService{}
}
