package util

import (
	"context"
	"fmt"

	"github.com/api-monolith-template/internal/constant"
	"github.com/api-monolith-template/internal/model/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func ContinueOrFatal(err error) {
	if err != nil {
		logrus.Fatal(err.Error())
	}
}

func GetUserIDFromContext(ctx context.Context) (*uuid.UUID, error) {
	userID, err := uuid.Parse(fmt.Sprintf("%s", ctx.Value(constant.UserID)))
	if err != nil {
		return nil, err
	}
	return &userID, nil
}

func GetTokenIDFromContext(ctx context.Context) (*uuid.UUID, error) {
	tokenID, err := uuid.Parse(fmt.Sprintf("%s", ctx.Value(constant.TokenID)))
	if err != nil {
		return nil, err
	}
	return &tokenID, nil
}

func HandleResponse(ctx *gin.Context, resp *response.BaseResponse, err error) {
	if err != nil {
		HandleError(ctx, err)
		ctx.Abort()
		return
	}
	// set default message
	if resp.Message == "" && resp.StatusCode < 300 && resp.StatusCode >= 200 {
		resp.Message = response.MessageOK
	}
	ctx.JSON(resp.StatusCode, resp)
}
