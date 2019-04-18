package msg

type Demo5Error struct {
	code    int
	message string
}

func New(code int) Demo5Error {
	msg, ok := m[code]

	if !ok {
		msg = "未知错误"
	}

	return Demo5Error{
		code:    code,
		message: msg,
	}
}

func (e Demo5Error) Error() string {
	return e.message
}

func (e Demo5Error) GetCode() int {
	return e.code
}
