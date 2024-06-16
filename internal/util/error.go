package util

import (
	"net/http"

	"github.com/api-monolith-template/internal/model/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleError(ctx *gin.Context, err error) {
	switch cErr := err.(type) {
	case *response.CustomError:
		ctx.JSON(cErr.StatusCode, cErr.Message)
	case validator.ValidationErrors:
		validationErr := processValidationErr(cErr)
		ctx.JSON(http.StatusInternalServerError, validationErr)
	default:
		ctx.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: "internal server error",
		})
	}
}

func processValidationErr(fieldErrs validator.ValidationErrors) *response.BaseResponse {
	var validationErrors []response.ValidationError

	for _, fieldError := range fieldErrs {
		validationError := response.ValidationError{
			Field:   fieldError.Field(),
			Tag:     fieldError.Tag(),
			Message: fieldError.Translate(UniversalTranslator.GetFallback()), // TODO: get lang from req header
		}
		validationErrors = append(validationErrors, validationError)
	}

	return &response.BaseResponse{
		StatusCode:       http.StatusBadRequest,
		Message:          "validation error",
		ValidationErrors: validationErrors,
	}
}
