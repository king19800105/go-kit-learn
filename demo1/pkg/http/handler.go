package http

import (
	"net/http"
	"github.com/king19800105/go-kit-learn/demo1/pkg/endpoint"
	"encoding/json"
	"context"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/king19800105/go-kit-learn/demo1/msg"
)

// 统一json返回格式
type responseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// 空json返回
type nop struct{}

// create的http请求设置，以及请求参数和响应参数的格式化
func makeCreateHandler(m *mux.Router, eps endpoint.Endpoints, options []kithttp.ServerOption) {
	m.Methods("POST", "GET").Path("/order/create").Handler(
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedMethods([]string{"POST", "GET"}),
		)(kithttp.NewServer(eps.CreateEndpoint, decodeCreateRequest, encodeCreateResponse, options...)))
}

// 解析请求参数
func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.CreateRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, msg.New(msg.JSON_FORMAT_FAILED)
	}

	return request, nil
}

// 编码响应
func encodeCreateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(kitendpoint.Failer); ok && nil != f.Failed() {
		return f.Failed()
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// 成功的响应
	json.NewEncoder(w).Encode(responseWrapper{
		Code:    msg.SUCCESS,
		Message: msg.GetMsg(msg.SUCCESS),
		Data:    response,
	})

	return nil
}

// 失败的统一格式化
func ErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	var code int
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(err2code(err))

	if err, ok := err.(msg.Demo1Error); ok {
		code = err.GetCode()
	} else {
		code = -1
	}

	json.NewEncoder(w).Encode(responseWrapper{
		Code:    code,
		Message: err.Error(),
		Data:    nop{},
	})
}

// 响应码设置
func err2code(err error) int {
	return http.StatusInternalServerError
}
