package constant

import (
	"net/http"

	"github.com/api-monolith-template/internal/model/response"
)

var (
	ErrInternalServerError = response.BaseResponse{
		ErrorCode:  "500_1",
		StatusCode: http.StatusInternalServerError,
		Message:    "internal server error",
	}
	ErrValidationError = response.BaseResponse{
		ErrorCode:  "400_1",
		StatusCode: http.StatusBadRequest,
		Message:    "validation error",
	}
)
