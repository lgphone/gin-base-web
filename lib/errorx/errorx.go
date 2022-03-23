package errorx

type AuthError struct {
	Code int
	msg  string
}

func (e AuthError) Error() string {
	if e.msg == "" {
		e.msg = "认证错误!"
	}
	return e.msg
}

func NewAuthError(msg string) *AuthError {
	return &AuthError{Code: 1001, msg: msg}
}

type ParamError struct {
	Code int
	msg  string
}

func (e ParamError) Error() string {
	if e.msg == "" {
		e.msg = "参数错误!"
	}
	return e.msg
}

func NewParamError(msg string) *ParamError {
	return &ParamError{Code: 1002, msg: msg}
}

type LogicError struct {
	Code int
	msg  string
}

func (e LogicError) Error() string {
	if e.msg == "" {
		e.msg = "逻辑错误!"
	}
	return e.msg
}

func NewLogicError(msg string) *LogicError {
	return &LogicError{Code: 1003, msg: msg}
}
