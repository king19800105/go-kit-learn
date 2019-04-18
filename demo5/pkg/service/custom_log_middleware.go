package service

import (
	"github.com/go-kit/kit/log"
	"context"
	"github.com/king19800105/go-kit-learn/demo5/pkg/entity"
)

// 自定义服务中间件
type Middleware func(OrderService) OrderService

//  日志中间件结构对象
type loggingMiddleware struct {
	logger log.Logger   // 需要使用的日志对象
	next   OrderService // 中间件的后续执行
}

// 创建中间件对象，由于中间件实现了OrderService方法，所以它也是OrderService对象
func LoggingMiddleware(logger log.Logger) Middleware {
	// log.NewNopLogger() 使用NopLog则表示取消日志功能
	return func(next OrderService) OrderService {
		return &loggingMiddleware{logger, next}
	}
}

// 中间件注入到Create中
func (log loggingMiddleware) Create(ctx context.Context, orderId string) (o entity.Order, err error) {
	// 最后执行
	defer func() {
		log.logger.Log("服务中间件：method", "Foo", "orderId", orderId, "o", o, "err", err)
	}()

	log.logger.Log("服务中间件：Create 方法开始执行前")
	o, err = log.next.Create(ctx, orderId)
	log.logger.Log("服务中间件：Create 方法执行完成后")

	return o, err
}
