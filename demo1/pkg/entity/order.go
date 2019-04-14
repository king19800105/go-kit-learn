package entity

// 返回一个订单创建成果的结构
type Order struct {
	Id     string `json:"orderId"`
	Source string `json:"source"`
	IsPay  int    `json:"isPay"`
}
