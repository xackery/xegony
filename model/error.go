package model

import (
	"github.com/pkg/errors"
)

//Error is used for Xegony Error handling
// swagger:response
type Error struct {
	ID         int64  `json:"id" db:"id"`                  //`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	URL        string `json:"url" db:"url"`                //`url` varchar(32) NOT NULL DEFAULT '',
	Scope      string `json:"scope" db:"scope"`            //`scope` varchar(32) NOT NULL DEFAULT '',
	Message    string `json:"message" db:"message"`        //`message` varchar(256) NOT NULL,
	Severity   int64  `json:"severity" db:"severity"`      //`severity` int(10) unsigned NOT NULL DEFAULT '0',
	CreateDate int64  `json:"createDate" db:"create_date"` //`create_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
}

// ErrNoContent is an error that is used when no content should be displayed
// swagger:response ErrNoContent
type ErrNoContent struct {
}

//Error wraps the error message to satisfy the error type interface
func (e *ErrNoContent) Error() string {
	return "No content provided"
}

//ErrDecodeBody is a failure to decode a request body
// swagger:response
type ErrDecodeBody struct {
}

//Error wraps the error message to satisfy the error type interface
func (e *ErrDecodeBody) Error() string {
	return "Failed to decode body"
}

//ErrInvalidArguments means arguments being passed in a request were invalid
// swagger:response
type ErrInvalidArguments struct {
}

//Error wraps the error message to satisfy the error type interface
func (e *ErrInvalidArguments) Error() string {
	return "Invalid arguments provided"
}

//ErrValidation has many errors represented as Key/Value pairs of Field:Description inside REasons
// swagger:response
type ErrValidation struct {
	Message string
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

//ErrPermission is a permission denied generic error
// swagger:response
type ErrPermission struct {
	Message string
}

//Error wraps the error message to satisfy the error type interface
func (e *ErrPermission) Error() string {
	return e.Message
}
