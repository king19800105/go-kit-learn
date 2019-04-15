package service

import (
	"github.com/king19800105/go-kit-learn/demo3/pkg/service"
	"github.com/king19800105/go-kit-learn/demo3/pkg/endpoint"
	"github.com/king19800105/go-kit-learn/demo3/pkg/http"
	kitendpoint "github.com/go-kit/kit/endpoint"
	nethttp "net/http"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"context"
)

// 创建服务
func createService(logger log.Logger, tracer opentracing.Tracer) nethttp.Handler {
	// 创建业务对象
	svc := service.New(registerServiceMiddleware(logger))
	// 创建端点对象
	eps := endpoint.New(svc, registerEndpointMiddleware(logger))
	// 设置http服务服务中间件
	options := defaultHttpOptions(logger, tracer)
	// 端点绑定到http服务上
	return http.NewHTTPHandler(eps, options)
}

// 自定义log，在下面使用
type MyLog struct {}

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
			kithttp.ServerBefore(setCtxBeforeRequest),
			kithttp.ServerAfter(handlerAfterResponse),
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
	duration := kitprometheus.NewSummaryFrom(prometheus.SummaryOpts{
		Help:      "Request duration in seconds.",
		Name:      "request_duration_seconds",
		Namespace: "example",
		Subsystem: "mytest",
	}, []string{"method", "success"})

	mw["Create"] = []kitendpoint.Middleware{
		endpoint.LoggingMiddleware(log.With(logger, "method", "Create")),
		endpoint.InstrumentingMiddleware(duration.With("method", "Create")),
	}
	// Add you endpoint middleware here

	return
}
