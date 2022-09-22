package errors

import (
	"fmt"
)

type Error interface {
	error
	GetCode() int
	GetMessage() string
}

type CustomError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

func (c CustomError) GetCode() int {
	return c.Code
}

func (c CustomError) GetMessage() string {
	return c.Message
}

func (c CustomError) Error() string {
	return fmt.Sprintf("[%v] %s", c.Code, c.Message)
}

func NewCustomHttpError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func NewCustomHttpErrorWithCode(code int, msg string, statusCode string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: fmt.Sprintf("%s with status code: [%v]", msg, statusCode),
	}
}

func NewSystemError(msg string) *CustomError {
	return &CustomError{
		Code:    SystemError,
		Message: msg,
	}
}

func NewSystemErrorWithCode(statusCode string) *CustomError {
	return &CustomError{
		Code:    SystemError,
		Message: fmt.Sprintf("system error with status_code:[%v]", statusCode),
	}
}

func NewBadRequestWithCode(statusCode string) *CustomError {
	return &CustomError{
		Code:    BadRequest,
		Message: fmt.Sprintf("bad request with status_code:[%v]", statusCode),
	}
}

func NewUnknownError(statusCode string) *CustomError {
	msg := fmt.Sprintf("unknown error with code:[%v]", statusCode)
	return NewSystemError(msg)
}
