package myerror

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Error struct {
	HttpCode   int
	ErrorCode  int
	Message    string
	InnerError error
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) ToHTTPError() *echo.HTTPError {
	return echo.NewHTTPError(e.HttpCode, e)
}

func NewBadRequestError(err error, message string, errCode int) *Error {
	return &Error{
		HttpCode:   http.StatusBadRequest,
		ErrorCode:  errCode,
		Message:    message,
		InnerError: err,
	}
}

func NewUnauthorizedError(err error, message string, errCode int) *Error {
	return &Error{
		HttpCode:   http.StatusUnauthorized,
		ErrorCode:  errCode,
		Message:    message,
		InnerError: err,
	}
}

func NewForbiddenError(err error, message string, errCode int) *Error {
	return &Error{
		HttpCode:   http.StatusForbidden,
		ErrorCode:  errCode,
		Message:    message,
		InnerError: err,
	}
}

func NewNotFoundError(err error, message string, errCode int) *Error {
	return &Error{
		HttpCode:   http.StatusNotFound,
		ErrorCode:  errCode,
		Message:    message,
		InnerError: err,
	}
}

func NewInternalServerError(err error, message string, errCode int) *Error {
	return &Error{
		HttpCode:   http.StatusInternalServerError,
		ErrorCode:  errCode,
		Message:    message,
		InnerError: err,
	}
}
