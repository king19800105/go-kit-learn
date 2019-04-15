package http

import (
	"github.com/king19800105/go-kit-learn/demo1/pkg/endpoint"
	"net/http"
	"github.com/gorilla/mux"
	kithttp "github.com/go-kit/kit/transport/http"
)

// 初始化监听的路由和处理的端点绑定
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]kithttp.ServerOption) http.Handler {
	m := mux.NewRouter()
	makeCreateHandler(m, endpoints, options["Create"])

	return m
}