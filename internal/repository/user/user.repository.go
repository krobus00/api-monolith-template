package user

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository() *Repository {
	return new(Repository)
}

func (r *Repository) WithGormDB(db *gorm.DB) *Repository {
	r.db = db
	return r
}
