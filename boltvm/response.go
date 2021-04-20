package boltvm

type VmStatus int

const (
	Normal VmStatus = iota
	Unknown
	NoBindRule
	NotAvailableAppchain
	DecodeFail
	EncodeFail
	InvalidIBTP
	ExistIBTPIndex
	WrongIBTPIndex
)

type Response struct {
	Code   VmStatus
	Result []byte
}

// Result returns normal result
func Success(data []byte) *Response {
	return &Response{
		Code:   Normal,
		Result: data,
	}
}

// Error returns error result that will cause
// vm call error, and this transaction will be invalid
func Error(msg string, status VmStatus) *Response {
	return &Response{
		Code:   status,
		Result: []byte(msg),
	}
}
