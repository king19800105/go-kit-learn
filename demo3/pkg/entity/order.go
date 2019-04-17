package entity

import "encoding/json"

// 返回一个订单创建成果的结构
type Order struct {
	Id     string `json:"orderId"`
	Source string `json:"source"`
	IsPay  int32  `json:"isPay"`
}

// 自动转换
func (t Order) String() string {
	b, err := json.Marshal(t)

	if err != nil {
		return "unsupported value type"
	}

	return string(b)
}
