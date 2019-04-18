package service

import (
	"github.com/king19800105/go-kit-learn/demo5/pkg/service"
	"github.com/king19800105/go-kit-learn/demo5/pkg/endpoint"
	"github.com/king19800105/go-kit-learn/demo5/pkg/http"
	kitendpoint "github.com/go-kit/kit/endpoint"
	nethttp "net/http"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"context"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	googlegrpc "google.golang.org/grpc"
	"github.com/juju/ratelimit"
	"time"
	"github.com/afex/hystrix-go/hystrix"
)

// 创建http服务
func createHttpService(logger log.Logger, tracer opentracing.Tracer) nethttp.Handler {
	// 设置基础服务
	eps := loadBaseService(logger)
	// 设置http服务服务中间件
	options := defaultHttpOptions(logger, tracer)
	// 端点绑定到http服务上
	return http.NewHTTPHandler(eps, options)
}

// 创建grpc服务
func createGRPCService(logger log.Logger, tracer opentracing.Tracer) *googlegrpc.Server {
	eps := loadBaseService(logger)
	options := defaultGrpcOptions(logger, tracer)

	return http.NewGRPCHandler(eps, options)
}

// 加载基础服务，以及中间件
func loadBaseService(logger log.Logger) endpoint.Endpoints {
	// 创建业务对象
	svc := service.New(registerServiceMiddleware(logger))
	// 创建端点对象
	return endpoint.New(svc, registerEndpointMiddleware(logger))
}

// 自定义log，在下面使用
type MyLog struct{}

func (l MyLog) Log(keyvals ...interface{}) error {
	fmt.Println(keyvals[0], keyvals[1])
	return nil
}

// HTTP服务中间件（服务的aop）
func defaultHttpOptions(logger log.Logger, tracer opentracing.Tracer) map[string][]kithttp.ServerOption {
	options := map[string][]kithttp.ServerOption{
		"Create": {
			// 加载统一服务响应
			kithttp.ServerErrorEncoder(http.ErrorEncoder),
			// 当请求或响应时发生的错误，使用日志捕获，可用MyLog对象来替换
			kithttp.ServerErrorLogger(logger),
			// 请求开始之前钩子
			kithttp.ServerBefore(setCtxBeforeRequest),
			// 请求结束之后钩子
			kithttp.ServerAfter(handlerAfterResponse),
		},
	}

	return options
}

// grpc中间件设置，和http请求中间件设置方式一样
func defaultGrpcOptions(logger log.Logger, tracer opentracing.Tracer) map[string][]kitgrpc.ServerOption {
	options := map[string][]kitgrpc.ServerOption{
		"Create": {
			// 自定义错误日志对象
			kitgrpc.ServerErrorLogger(MyLog{}),
		},
	}

	return options
}

// 注册before事件
func setCtxBeforeRequest(ctx context.Context, req *nethttp.Request) context.Context {
	fmt.Println("在http请求之前，设置heard头，以及设置ctx对象，在整个生命周期都有效")
	req.Header.Set("Authorization", "abcdefg")
	ctx = context.WithValue(ctx, "ctxSet", "myValue")

	return ctx
}

// 注册after事件
func handlerAfterResponse(ctx context.Context, res nethttp.ResponseWriter) context.Context {
	val := ctx.Value("ctxSet")
	fmt.Println("在http请求完成之后执行，得到ctx中的值：" + val.(string))

	return ctx
}

// 服务中间件注册
func registerServiceMiddleware(logger log.Logger) (mw []service.Middleware) {
	mw = []service.Middleware{}
	// 加载自定义日志对象
	mw = append(mw, service.LoggingMiddleware(logger))
	// Append your middleware here

	return
}

// 端点中间件注册
func registerEndpointMiddleware(logger log.Logger) (mw map[string][]kitendpoint.Middleware) {
	mw = map[string][]kitendpoint.Middleware{}
	// 监控中间件
	fieldKeys := []string{"method"}
	requestCount := kitprometheus.NewCounterFrom(prometheus.CounterOpts{
		Namespace: "demo5",
		Subsystem: "order_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(prometheus.SummaryOpts{
		Namespace: "demo5",
		Subsystem: "order_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	// 限流中间件，1秒钟接口并发1次
	rlbucket := ratelimit.NewBucket(1*time.Second, 1)
	// 断路器中间件
	command := "Create Request"
	hystrix.ConfigureCommand(command, hystrix.CommandConfig{Timeout: 1000})
	mw["Create"] = []kitendpoint.Middleware{
		endpoint.LoggingMiddleware(log.With(logger, "method", "Create")),
		endpoint.InstrumentingMiddleware(requestCount, requestLatency),
		endpoint.RateLimiterMiddleware(rlbucket),
		endpoint.HystrixMiddleware(command, "Service currently unavailable", logger),
	}
	// Add you endpoint middleware here

	return
}
