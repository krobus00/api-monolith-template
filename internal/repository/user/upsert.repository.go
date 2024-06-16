package user

import (
	"context"

	"github.com/api-monolith-template/internal/model/entity"
	"github.com/api-monolith-template/internal/util"
	"github.com/sirupsen/logrus"
)

func (r *Repository) Upsert(ctx context.Context, user *entity.User) error {
	logger := util.NewDefaultLogger(ctx).WithFields(logrus.Fields{
		"user": util.Dump(user),
	})
	tx := util.GetTxFromContext(ctx, r.db)

	err := tx.Save(&user).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
