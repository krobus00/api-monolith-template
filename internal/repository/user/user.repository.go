package user

import (
	"github.com/api-monolith-template/internal/model/contract"
	"gorm.io/gorm"
)

type Repository struct {
	db        *gorm.DB
	cacheRepo contract.CacheRepository
}

func NewRepository() *Repository {
	return new(Repository)
}

func (r *Repository) WithGormDB(db *gorm.DB) *Repository {
	r.db = db
	return r
}

func (r *Repository) WithCacheRepository(repo contract.CacheRepository) *Repository {
	r.cacheRepo = repo
	return r
}
