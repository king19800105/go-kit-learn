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
)

// create的http请求设置，以及请求参数和响应参数的格式化
func makeCreateHandler(m *mux.Router, eps endpoint.Endpoints) {
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
		return nil, msg.GetErr(msg.JSON_FORMAT_ILLEGAL)
	}

	return request, nil
}

// 编码响应
func encodeCreateResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
