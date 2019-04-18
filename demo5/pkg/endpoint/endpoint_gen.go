package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/king19800105/go-kit-learn/demo5/pkg/service"
)

type Endpoints struct {
	CreateEndpoint endpoint.Endpoint
}

// 封装打包多个端点，并使用包装模式加载中间件
func New(s service.OrderService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		CreateEndpoint: makeCreateEndpoint(s),
	}

	for _, m := range mdw["Create"] {
		eps.CreateEndpoint = m(eps.CreateEndpoint)
	}

	return eps
}
