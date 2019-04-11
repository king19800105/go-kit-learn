package msg

import "errors"

const (
	SUCCESS             = 0
	JSON_FORMAT_ILLEGAL = 1000
	ORDER_NO_EMPTY      = 1001
)

var tipsMsg = map[int]string{
	SUCCESS: "success",
}

var errMsg = map[int]error{
	JSON_FORMAT_ILLEGAL: errors.New("json参数格式不正确"),
	ORDER_NO_EMPTY:      errors.New("订单编号不能为空"),
}

func GetErr(code int) error {
	return errMsg[code]
}

func GetCodeErr(code int) (int, error) {
	return code, GetErr(code)
}
