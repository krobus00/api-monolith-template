package constant

import (
	"net/http"

	"github.com/api-monolith-template/internal/model/response"
)

var (
	ErrInternalServerError = response.CustomError{
		Code:       "INTERNAL_SERVER_ERROR",
		StatusCode: http.StatusInternalServerError,
		Message:    "internal server error",
	}
	ErrValidationError = response.CustomError{
		Code:       "VALIDATION_ERROR",
		StatusCode: http.StatusBadRequest,
		Message:    "validation error",
	}
	ErrPasswordNotMatch = response.CustomError{
		Code:       "PASSWORD_NOT_MATCH",
		StatusCode: http.StatusBadRequest,
		Message:    "password not match",
	}
	ErrUserNotFound = response.CustomError{
		Code:       "USER_NOT_FOUND",
		StatusCode: http.StatusNotFound,
		Message:    "user not found",
	}
	ErrInvalidToken = response.CustomError{
		Code:       "INVALID_TOKEN",
		StatusCode: http.StatusUnauthorized,
		Message:    "invalid token",
	}
	ErrTokenExpired = response.CustomError{
		Code:       "TOKEN_EXPIRED",
		StatusCode: http.StatusUnauthorized,
		Message:    "token expired",
	}
	ErrEndpointNotFound = response.CustomError{
		Code:       "ENDPOINT_NOT_FOUND",
		StatusCode: http.StatusNotFound,
		Message:    "endpoint not found",
	}
)
