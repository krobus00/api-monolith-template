package user

import (
	"context"

	"github.com/api-monolith-template/internal/model/entity"
	"github.com/api-monolith-template/internal/util"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (r *Repository) FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	logger := util.NewDefaultLogger(ctx).WithFields(logrus.Fields{
		"id": id,
	})

	result := new(entity.User)

	tx := util.GetTxFromContext(ctx, r.db)

	err := tx.Where("id = ?", id).First(&result).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return result, nil
}
