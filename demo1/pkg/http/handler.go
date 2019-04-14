package http

import (
	"net/http"
	"github.com/king19800105/go-kit-learn/demo1/pkg/endpoint"
	"encoding/json"
	"context"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/king19800105/go-kit-learn/internal/demo1/msg"
	kitendpoint "github.com/go-kit/kit/endpoint"
)

type responseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type nop struct {

}

// create的http请求设置，以及请求参数和响应参数的格式化
func makeCreateHandler(m *mux.Router, eps endpoint.Endpoints) {
	// todo... 测试Get方式是否也可以用json方式
	m.Methods("POST").Path("/order/create").Handler(
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedMethods([]string{"POST"}),
		)(kithttp.NewServer(eps.CreateEndpoint, decodeCreateRequest, encodeCreateResponse)))
}

// 解析请求参数
func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.CreateRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, msg.GetErr(msg.JSON_FORMAT_FAILED)
	}

	return request, nil
}

// 编码响应
func encodeCreateResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if f, ok := response.(kitendpoint.Failer); ok && nil != f.Failed() {
		ErrorEncoder(w, f.Failed(), msg.ORDER_CREATE_FAILED)
		return nil
	}

	SuccessEncoder(w, response)

	return nil
}

func SuccessEncoder(w http.ResponseWriter, res interface{}) {
	json.NewEncoder(w).Encode(responseWrapper{
		Code:    msg.SUCCESS,
		Message: msg.GetSuccess(),
		Data:    res,
	})
}

func ErrorEncoder(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(responseWrapper{
		Code: code,
		Message: err.Error(),
		Data: nop{},
	})
}

// 响应码设置
func err2code(err error) int {
	return http.StatusInternalServerError
}
