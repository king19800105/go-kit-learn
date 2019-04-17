package msg

const (
	SUCCESS             = 0
	ORDER_CREATE_FAILED = 1000
	JSON_FORMAT_FAILED  = 1001
	ORDER_NO_EMPTY      = 1002
)

var m = map[int]string{
	SUCCESS:             "success",
	ORDER_CREATE_FAILED: "创建订单失败",
	JSON_FORMAT_FAILED:  "请求参数不是一个合法的json格式",
	ORDER_NO_EMPTY:      "订单编号不能为空",
}

func GetMsg(code int) string {
	return m[code]
}
