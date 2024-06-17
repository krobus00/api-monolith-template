package util

import (
	"github.com/api-monolith-template/internal/constant"
	"github.com/api-monolith-template/internal/model/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleError(ctx *gin.Context, err error) {
	switch cErr := err.(type) {
	case response.CustomError:
		ctx.JSON(cErr.StatusCode, cErr.ToResponse())
	case validator.ValidationErrors:
		validationErr := processValidationErr(cErr)
		ctx.JSON(validationErr.StatusCode, validationErr)
	default:
		internalServerErr := constant.ErrInternalServerError.ToResponse()
		ctx.JSON(internalServerErr.StatusCode, internalServerErr)
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

	resp := constant.ErrValidationError.ToResponse()
	resp.ValidationErrors = validationErrors

	return &resp
}
