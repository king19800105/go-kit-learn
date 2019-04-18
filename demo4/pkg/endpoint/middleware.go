package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"time"
	"github.com/juju/ratelimit"
	"errors"
	"github.com/afex/hystrix-go/hystrix"
)

// 端点中间件，和服务中间件不同，前者是在服务执行前，执行后调用。后者是在端点调用前，端点调用后执行

// 监控中间件
func InstrumentingMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				lvs := []string{"method", "Create"}
				requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
				requestCount.With(lvs...).Add(1)
			}(time.Now())
			return next(ctx, request)
		}
	}
}

// 端点的日志中间件
func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Log("端点中间件：error", err, "执行时间：", time.Since(begin))
			}(time.Now())
			return next(ctx, request)
		}
	}
}

// 限流中间件
func RateLimiterMiddleware(tb *ratelimit.Bucket) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			if 0 == tb.TakeAvailable(1) {
				return nil, errors.New("Rate Limit Exceed")
			}

			return next(ctx, request)
		}
	}
}

// 断路器中间件
func HystrixMiddleware(commandName string, fallbackMsg string, logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			var resp interface{}

			if err := hystrix.Do(commandName, func() error {
				resp, err = next(ctx, request)
				return err
			}, func(e error) error {
				logger.Log("fallbackErrorDesc", err.Error())
				resp = struct {
					Fallback string `json:"fallback"`
				}{
					fallbackMsg,
				}

				return nil
			}); nil != err {
				return nil, err
			}

			return resp, nil
		}
	}
}
