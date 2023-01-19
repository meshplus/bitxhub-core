package boltvm

import (
	"fmt"
)

type BxhError struct {
	Code ErrorCode
	Msg  ErrorMsg
}

func BError(code ErrorCode, msg string) *BxhError {
	return &BxhError{
		Code: code,
		Msg:  ErrorMsg(msg),
	}
}

type Response struct {
	Ok     bool
	Result []byte
}

type ErrorCode string
type ErrorMsg string

func (be BxhError) Error() string {
	return fmt.Sprintf("%s:%s", be.Code, be.Msg)
}

func (be BxhError) IsInternal() bool {
	return be.Code[0] == '2'
}

// Result returns normal result
func Success(data []byte) *Response {
	return &Response{
		Ok:     true,
		Result: data,
	}
}

// Error returns error result that will cause
// vm call error, and this transaction will be invalid
func Error(code ErrorCode, msg string) *Response {
	be := BxhError{
		Code: code,
		Msg:  ErrorMsg(msg),
	}

	return &Response{
		Ok:     false,
		Result: []byte(be.Error()),
	}
}

func ResponseWrapper(ok bool, data []byte) *Response {
	if ok {
		return Success(data)
	}
	return Error(OtherInternalErrCode, string(data))
}
