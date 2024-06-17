package contract

import (
	"context"

	"github.com/api-monolith-template/internal/model/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	FindByIdentifier(ctx context.Context, identifier string) (*entity.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	Upsert(ctx context.Context, user *entity.User) error
}
