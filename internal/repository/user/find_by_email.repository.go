package user

import (
	"context"

	"github.com/api-monolith-template/internal/model/entity"
	"github.com/api-monolith-template/internal/util"
	"github.com/sirupsen/logrus"
)

func (r *Repository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	logger := util.NewDefaultLogger(ctx).WithFields(logrus.Fields{
		"email": email,
	})

	result := new(entity.User)

	tx := util.GetTxFromContext(ctx, r.db)

	err := tx.Where("email = ?", email).First(&result).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return result, nil
}
