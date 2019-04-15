package http

import (
	"github.com/king19800105/go-kit-learn/demo2/pkg/endpoint"
	"net/http"
	"github.com/gorilla/mux"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 初始化监听的路由和处理的端点绑定
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]kithttp.ServerOption) http.Handler {
	m := mux.NewRouter()
	makeCreateHandler(m, endpoints, options["Create"])
	initMetricsEndpoint(m)

	return m
}

// 监控追踪注册
func initMetricsEndpoint(m *mux.Router) {
	m.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
}