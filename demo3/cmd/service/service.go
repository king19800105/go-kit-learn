package service

import (
	nethttp "net/http"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/go-kit/kit/log"
	"os"
)


// 服务运行
func Run() {
	// log.NewNopLogger() 替换下面的log对象后，取消日志
	logger := log.NewLogfmtLogger(os.Stderr)
	tracer := opentracing.GlobalTracer()
	// 添加日志相关组件
	httpHandler := createService(logger, tracer)
	logger.Log("demo3服务启动，服务地址: ", "127.0.0.1:8088")
	err := nethttp.ListenAndServe(":8088", httpHandler)

	if nil != err {
		fmt.Println(err)
	}
}
