package response

import (
	"time"

	"github.com/google/uuid"
)

type AuthResp struct {
	AccessToken           string    `json:"accessToken"`
	AccessTokenExpiredAt  time.Time `json:"accessTokenExpiredAt"`
	RefreshToken          string    `json:"refreshToken"`
	RefreshTokenExpiredAt time.Time `json:"refreshTokenExpiredAt"`
}

type AuthInfoResp struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Level     string    `json:"level"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
