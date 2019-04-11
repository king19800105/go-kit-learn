package service

import (
	"github.com/king19800105/go-kit-learn/demo1/pkg/service"
	"github.com/king19800105/go-kit-learn/demo1/pkg/endpoint"
	"github.com/king19800105/go-kit-learn/demo1/pkg/http"
	nethttp "net/http"
)

// 创建服务
func createService() nethttp.Handler {
	// 创建业务对象
	svc := service.New(nil)
	// 创建端点对象
	eps := endpoint.New(svc, nil)
	// 端点绑定到http服务上
	return http.NewHTTPHandler(eps)
}