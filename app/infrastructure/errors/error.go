package errors

import "fmt"

type ErrorStackFlow struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewErrorStackFlow(code int, Message string) *ErrorStackFlow {
	return &ErrorStackFlow{
		Code:    code,
		Message: Message,
	}
}

func (e *ErrorStackFlow) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

type ErrorStackFlows []ErrorStackFlow

var (
	ErrorResponse = ErrorStackFlow{Code: 400, Message: "response failed"}
)
