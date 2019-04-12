package http

import (
	"github.com/king19800105/go-kit-learn/demo1/pkg/endpoint"
	"net/http"
	"github.com/gorilla/mux"
)

// 初始化监听的路由和处理的端点绑定
func NewHTTPHandler(endpoints endpoint.Endpoints) http.Handler {
	m := mux.NewRouter()
	makeCreateHandler(m, endpoints)

	return m
}