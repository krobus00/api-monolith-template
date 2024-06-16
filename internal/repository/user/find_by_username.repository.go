package user

import (
	"context"

	"github.com/api-monolith-template/internal/model/entity"
	"github.com/api-monolith-template/internal/util"
	"github.com/sirupsen/logrus"
)

func (r *Repository) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	logger := util.NewDefaultLogger(ctx).WithFields(logrus.Fields{
		"username": username,
	})

	result := new(entity.User)

	tx := util.GetTxFromContext(ctx, r.db)

	err := tx.Where("username = ?", username).First(&result).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return result, nil
}
