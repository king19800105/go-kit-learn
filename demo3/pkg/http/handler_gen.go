package http

import (
	"github.com/king19800105/go-kit-learn/demo3/pkg/endpoint"
	"net/http"
	"github.com/gorilla/mux"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/king19800105/go-kit-learn/demo3/pkg/grpc/pb"
	"github.com/king19800105/go-kit-learn/demo3/pkg/grpc"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	googlegrpc "google.golang.org/grpc"
)

// 初始化监听的路由和处理的端点绑定
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]kithttp.ServerOption) http.Handler {
	m := mux.NewRouter()
	makeCreateHandler(m, endpoints, options["Create"])
	initMetricsEndpoint(m)

	return m
}

// 初始化GRPC服务绑定到端点
func NewGRPCHandler(endpoints endpoint.Endpoints, options map[string][]kitgrpc.ServerOption) *googlegrpc.Server {
	grpcHandler := grpc.NewGRPCServer(endpoints, options)
	baseServer := googlegrpc.NewServer()
	pb.RegisterOrderServer(baseServer, grpcHandler)

	return baseServer
}

// 监控追踪注册
func initMetricsEndpoint(m *mux.Router) {
	m.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
}
