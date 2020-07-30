package errors

import (
	"fmt"
)

type baseError struct {
	code int
	msg  string
}

func newBaseError(code int, err error) baseError {
	var msg string
	if err != nil {
		msg = err.Error()
	}

	if msg != "" {
		msg = fmt.Sprintf("%s: %s", getErrorMessagePrefix(code), msg)
	} else {
		msg = getErrorMessagePrefix(code)
	}

	return baseError{code, msg}
}

func newStateMachineCustomError(code int, err error, currentState string, destinationState string) baseError {
	var msg string
	if err != nil {
		msg = err.Error()
	}

	if msg != "" {
		msg = fmt.Sprintf("%s: %s", getErrorMessagePrefix(code), msg)
	} else {
		fmt.Println(code)
		msg = getErrorMessagePrefix(code)

		msg = fmt.Sprintf(msg, currentState, destinationState)
	}

	return baseError{code, msg}
}

func (e baseError) Error() string {
	if e.msg != "" {
		return fmt.Sprintf("(%d) %s", int(e.code), e.msg)
	} else {
		return fmt.Sprintf("(%d) %s", int(e.code), getErrorMessagePrefix(e.code))
	}
}

func (e baseError) ExtractJsonRpsErrorParams() (int, string) {
	return e.code, e.msg
}

func ExtractJsonRpsErrorParams(err error) (int, string, bool) {
	if baseErr, ok := err.(jsonRpsErrorParamsExtractor); ok {
		code, msg := baseErr.ExtractJsonRpsErrorParams()
		return code, msg, true
	} else {
		return 0, "", false
	}
}
