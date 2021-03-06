package utils

import (
	"errors"
	"net/http"
)

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal server error")

	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("your requested item was not found")

	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("your item already exist")

	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("given param is not valid")

	// Mongo Not Found
	ErrMongoNotFound = errors.New("mongo: no documents in result")

	// Invalid date of birthday
	ErrInvalidDateOfBirthday = errors.New("invalid date of birhday")
)

type ResponseError struct {
	Message string `json:"message"`
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

//BuildResponseFromError --
func BuildResponseFromError(err error) ResponseError {
	return ResponseError{
		Message: err.Error(),
	}
}
