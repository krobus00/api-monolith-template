package util

import (
	"context"
	"fmt"

	"github.com/api-monolith-template/internal/constant"
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
