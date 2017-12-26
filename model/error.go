package model

import (
	"github.com/pkg/errors"
)

type ErrNoContent struct {
}

func (e *ErrNoContent) Error() string {
	return "No content provided"
}

type ErrDecodeBody struct {
}

func (e *ErrDecodeBody) Error() string {
	return "Failed to decode body"
}

type ErrInvalidArguments struct {
}

func (e *ErrInvalidArguments) Error() string {
	return "Invalid arguments provided"
}

type ErrValidation struct {
	Message string
	Reasons map[string]string
}

func (e *ErrValidation) Error() string {
	return e.Message
}

type StackTracer interface {
	StackTrace() errors.StackTrace
}

type ErrPermission struct {
	Message string
}

func (e *ErrPermission) Error() string {
	return e.Message
}
