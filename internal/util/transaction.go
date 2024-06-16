package util

import (
	"context"

	"github.com/api-monolith-template/internal/constant"
	"gorm.io/gorm"
)

func GetTxFromContext(ctx context.Context, defaultTx *gorm.DB) *gorm.DB {
	txVal := ctx.Value(constant.DB)
	tx, ok := txVal.(*gorm.DB)
	if !ok {
		return defaultTx.WithContext(ctx)
	}
	return tx
}
