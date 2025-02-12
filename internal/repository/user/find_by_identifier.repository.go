package user

import (
	"context"
	"errors"

	"github.com/api-monolith-template/internal/model/cachekey"
	"github.com/api-monolith-template/internal/model/entity"
	"github.com/api-monolith-template/internal/util"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (r *Repository) FindByIdentifier(ctx context.Context, identifier string) (*entity.User, error) {
	logger := util.NewDefaultLogger(ctx).WithFields(logrus.Fields{
		"identifier": identifier,
	})

	result := new(entity.User)
	cacheKey := cachekey.NewUserByIdentifierCacheKey(identifier)
	err := r.cacheRepo.GetOrSetCache(ctx, cacheKey, &result, func(ctx context.Context) (any, error) {
		tx := util.GetTxFromContext(ctx, r.db)
		err := tx.Where("email = ? OR username = ?", identifier, identifier).First(&result).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
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
