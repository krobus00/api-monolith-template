package user

import (
	"context"

	"github.com/api-monolith-template/internal/model/cachekey"
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
	cacheKey := cachekey.NewUserByIDCacheKey(id.String())
	err := r.cacheRepo.GetOrSetCache(ctx, cacheKey, result, func(ctx context.Context) (any, error) {
		result := new(entity.User)
		tx := util.GetTxFromContext(ctx, r.db)
		err := tx.Where("id = ? ", id).First(&result).Error
		if err != nil {
			return nil, err
		}

		return result, nil
	})

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return result, nil
}
