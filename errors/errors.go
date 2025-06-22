package errors

import "errors"

type ErrorInfo struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

const (
	ErrMissingParam   = 101
	ErrJsonInvalid    = 102
	ErrInvalidInput   = 103
	ErrFileMissing    = 104
	ErrURLMissing     = 105
	ErrTechnicalissue = 110
)

var Errors = map[int]error{
	ErrMissingParam:   errors.New("Parameter is missing"),
	ErrJsonInvalid:    errors.New("Json is invalid"),
	ErrInvalidInput:   errors.New("Input is invalid"),
	ErrFileMissing:    errors.New("The file is either missing or the format of submitted file is invalid"),
	ErrURLMissing:     errors.New("The url is either missing or is not accessible"),
	ErrTechnicalissue: errors.New("Internal server error"),
}
