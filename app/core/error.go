package core

import (
	"gorm.io/gorm"
	"net/http"
)

type DetailedError interface {
	error
	Code() int
}

type Error struct {
	error
	message   string
	errorCode int
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) Code() int {
	return e.errorCode
}

func NewSimpleError(message string) DetailedError {
	return &Error{message: message}
}

func NewDatabaseError(err error) DetailedError {
	switch err {
	case gorm.ErrRecordNotFound:
		return NewNotFoundError(err.Error())
	case gorm.ErrDuplicatedKey:
		return NewUnprocessableEntityError(err.Error())
	}
	return NewSimpleError(err.Error())
}

func NewNotFoundError(message string) DetailedError {
	return &Error{errorCode: http.StatusNotFound, message: message}
}

func NewBadRequestError(message string) DetailedError {
	return &Error{errorCode: http.StatusBadRequest, message: message}
}

func NewUnauthorizedError(message string) DetailedError {
	return &Error{errorCode: http.StatusUnauthorized, message: message}
}

func NewUnprocessableEntityError(message string) DetailedError {
	return &Error{errorCode: http.StatusUnprocessableEntity, message: message}
}
