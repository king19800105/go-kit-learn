package service

import (
	"context"
	"github.com/king19800105/go-kit-learn/demo2/pkg/entity"
	"fmt"
	"github.com/king19800105/go-kit-learn/internal/demo2/msg"
)

// 服务抽象
type OrderService interface {
	Create(ctx context.Context, orderId string) (entity.Order, error)
}

// 订单结构
type baseOrderService struct{}


// 创建订单
func (os baseOrderService) Create(ctx context.Context, orderId string) (o entity.Order, err error) {
	if "" == orderId {
		return o, msg.New(msg.ORDER_NO_EMPTY)
	}
	val := ctx.Value("ctxSet")
	fmt.Println("在服务中获取http请求之前设置的ctx值：" + val.(string))
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
