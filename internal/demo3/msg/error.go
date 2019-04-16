package msg

type Demo3Error struct {
	code    int
	message string
}

func New(code int) Demo3Error {
	msg, ok := m[code]

	if !ok {
		msg = "未知错误"
	}

	return Demo3Error{
		code:    code,
		message: msg,
	}
}

func (e Demo3Error) Error() string {
	return e.message
}

func (e Demo3Error) GetCode() int {
	return e.code
}
