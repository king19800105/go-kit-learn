package msg

type Demo4Error struct {
	code    int
	message string
}

func New(code int) Demo4Error {
	msg, ok := m[code]

	if !ok {
		msg = "未知错误"
	}

	return Demo4Error{
		code:    code,
		message: msg,
	}
}

func (e Demo4Error) Error() string {
	return e.message
}

func (e Demo4Error) GetCode() int {
	return e.code
}
