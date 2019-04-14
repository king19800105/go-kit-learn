package msg

import "errors"

const (
	SUCCESS             = 0
	ORDER_CREATE_FAILED = 1000
	JSON_FORMAT_FAILED  = 1001
	ORDER_NO_EMPTY      = 1002
)

var errMsg = map[int]error{
	JSON_FORMAT_FAILED: errors.New("请求参数不是一个合法的json格式"),
	ORDER_NO_EMPTY:     errors.New("订单编号不能为空"),
}

func GetErr(code int) error {
	return errMsg[code]
}

func GetSuccess() string {
	return "success"
}
