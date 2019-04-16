package msg

type Demo1Error struct {
	code    int
	message string
}

func New(code int) Demo1Error {
	msg, ok := m[code]

	if !ok {
		msg = "未知错误"
	}

	return Demo1Error{
		code:    code,
		message: msg,
	}
}

func (e Demo1Error) Error() string {
	return e.message
}

func (e Demo1Error) GetCode() int {
	return e.code
}
