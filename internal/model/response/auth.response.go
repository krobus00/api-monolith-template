package response

import (
	"github.com/google/uuid"
)

type AuthResp struct {
	AccessToken           string `json:"accessToken"`
	AccessTokenExpiredAt  string `json:"accessTokenExpiredAt"`
	RefreshToken          string `json:"refreshToken"`
	RefreshTokenExpiredAt string `json:"refreshTokenExpiredAt"`
}

type AuthInfoResp struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Level     string    `json:"level"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
}
