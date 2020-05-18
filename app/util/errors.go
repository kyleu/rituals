package util

import (
	"context"
	"fmt"
	"strings"

	"emperror.dev/errors"

	"logur.dev/logur"
)

type AppErrorHandler struct {
	Logger logur.Logger
}

func (a *AppErrorHandler) Handle(err error) {
	if err != nil {
		a.Logger.Error(fmt.Sprintf("general error: %+v", err))
	}
}
func (*AppErrorHandler) HandleContext(_ context.Context, _ error) {}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

type unwrappable interface {
	Unwrap() error
}

type errorFrame struct {
	Key string
	Loc string
}

type ErrorDetail struct {
	Message    string
	StackTrace errors.StackTrace
	Cause      *ErrorDetail
}

func GetErrorDetail(e error) *ErrorDetail {
	var stack errors.StackTrace = nil
	t, ok := e.(stackTracer)
	if ok {
		stack = t.StackTrace()
	}

	var cause *ErrorDetail = nil
	u, ok := e.(unwrappable)
	if ok {
		cause = GetErrorDetail(u.Unwrap())
	}

	return &ErrorDetail{
		Message:    e.Error(),
		StackTrace: stack,
		Cause:      cause,
	}
}

func TraceDetail(trace errors.StackTrace) []errorFrame {
	s := fmt.Sprintf("%+v", trace)
	lines := strings.Split(s, "\n")
	validLines := make([]string, 0)
	for _, line := range lines {
		l := strings.TrimSpace(line)
		if len(l) > 0 {
			validLines = append(validLines, l)
		}
	}
	ret := make([]errorFrame, 0)
	for i := 0; i < len(validLines)-1; i += 2 {
		f := errorFrame{Key: validLines[i], Loc: validLines[i+1]}
		ret = append(ret, f)
	}
	return ret
}
