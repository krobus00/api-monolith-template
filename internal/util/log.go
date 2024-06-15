package util

import (
	"context"

	"github.com/api-monolith-template/internal/constant"
	"github.com/sirupsen/logrus"
)

func NewDefaultLogger(ctx context.Context) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"requestID": ctx.Value(constant.RequestID),
		"userID":    ctx.Value(constant.UserID),
		"userType":  ctx.Value(constant.UserType),
	})
}
