package constant

import (
	"net/http"

	"github.com/api-monolith-template/internal/model/response"
)

var (
	ErrInternalServerError = response.CustomError{
		Code:       "500_1",
		StatusCode: http.StatusInternalServerError,
		Message:    "internal server error",
	}
	ErrValidationError = response.CustomError{
		Code:       "400_1",
		StatusCode: http.StatusBadRequest,
		Message:    "validation error",
	}
	ErrPasswordNotMatch = response.CustomError{
		Code:       "400_2",
		StatusCode: http.StatusBadRequest,
		Message:    "password not match",
	}
	ErrUserNotFound = response.CustomError{
		Code:       "400_3",
		StatusCode: http.StatusBadRequest,
		Message:    "user not found",
	}
)
