package model

import (
	"github.com/pkg/errors"
)

//Error is used for Xegony Error handling
// swagger:model
type Error struct {
	ID         int64  `json:"id" db:"id"`                  //`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	URL        string `json:"url" db:"url"`                //`url` varchar(32) NOT NULL DEFAULT '',
	Scope      string `json:"scope" db:"scope"`            //`scope` varchar(32) NOT NULL DEFAULT '',
	Message    string `json:"message" db:"message"`        //`message` varchar(256) NOT NULL,
	Severity   int64  `json:"severity" db:"severity"`      //`severity` int(10) unsigned NOT NULL DEFAULT '0',
	CreateDate int64  `json:"createDate" db:"create_date"` //`create_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
}

// ErrNoContent is an is used when no content should be displayed
//
// No content, or nothing changed.
// swagger:model ErrNoContent
type ErrNoContent struct {
}

func (e *ErrNoContent) Error() string {
	return ""
}

// ErrInternal is an internal server error
// swagger:model
type ErrInternal struct {
	//example: internal server error
	Message string
}

func (e *ErrInternal) Error() string {
	return e.Message
}

//ErrDecodeBody is a failure to decode a request body
// swagger:model
type ErrDecodeBody struct {
	//example: failed to decode body
	Message string
}

//Error wraps the error message to satisfy the error type interface
func (e *ErrDecodeBody) Error() string {
	return e.Message
}

//ErrInvalidArguments means arguments being passed in a request were invalid
// swagger:model
type ErrInvalidArguments struct {
}

//Error wraps the error message to satisfy the error type interface
func (e *ErrInvalidArguments) Error() string {
	return "Invalid arguments provided"
}

//ErrRedirect means arguments being passed in a request were invalid
// swagger:model
type ErrRedirect struct {
	// example: http//google.signon.com
	Message string
}

//Error wraps the error message to satisfy the error type interface
func (e *ErrRedirect) Error() string {
	return "Redirecting you to new location"
}

//ErrValidation has many errors represented as Key/Value pairs of Field:Description inside REasons
// swagger:model
type ErrValidation struct {
	// example: failed to validate
	Message string
	// example: test
	Reasons map[string]string
}

//Error wraps the error message to satisfy the error type interface
func (e *ErrValidation) Error() string {
	return e.Message
}

//StackTracer is an interface for stack trace error handling
type StackTracer interface {
	StackTrace() errors.StackTrace
}

// ErrPermission is a permission denied generic error
// swagger:model
type ErrPermission struct {
	// example: permission denied
	Message string
}

//Error wraps the error message to satisfy the error type interface
func (e *ErrPermission) Error() string {
	return e.Message
}
