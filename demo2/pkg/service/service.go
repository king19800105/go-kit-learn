package service

import (
	"context"
	"log"
	"github.com/king19800105/go-kit-learn/internal/demo1/msg"
)

// 服务抽象
type OrderService interface {
	Create(ctx context.Context, orderId string) (code int, err error)
}

// 订单结构
type orderService struct{}

// 创建订单
func (order orderService) Create(ctx context.Context, orderId string) (code int, err error) {
	// 获取ctx对象
	val := ctx.Value("ID")
	log.Println(val)

	if "" == orderId {
		return msg.GetCodeErr(msg.JSON_FORMAT_ILLEGAL)
	}

	return 0, nil
}

// 服务对象实例化，并且组装中间件
func New(middleware []Middleware) OrderService {
	var svc = getServiceInstance()

	for _, m := range middleware {
		svc = m(svc)
	}

	return svc
}

// 获取当前实例
func getServiceInstance() OrderService {
	return &orderService{}
}
