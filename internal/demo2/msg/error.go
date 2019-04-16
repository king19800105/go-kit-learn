package msg

type Demo2Error struct {
	code    int
	message string
}

func New(code int) Demo2Error {
	msg, ok := m[code]

	if !ok {
		msg = "未知错误"
	}

	return Demo2Error{
		code:    code,
		message: msg,
	}
}

func (e Demo2Error) Error() string {
	return e.message
}

func (e Demo2Error) GetCode() int {
	return e.code
}
