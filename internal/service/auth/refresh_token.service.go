package auth

import (
	"context"
	"time"

	"github.com/api-monolith-template/internal/config"
	"github.com/api-monolith-template/internal/model/cachekey"
	"github.com/api-monolith-template/internal/model/request"
	"github.com/api-monolith-template/internal/model/response"
	"github.com/api-monolith-template/internal/util"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (s *Service) RefreshToken(ctx context.Context, req *request.AuthRefreshReq) (*response.BaseResponse, error) {
	logger := util.NewDefaultLogger(ctx).WithFields(logrus.Fields{
		"req": req,
	})

	accessTokenID := uuid.New()
	accessToken, accessExpiredAt, err := util.GenerateToken(config.Env.Token.AccessTokenSecret, req.UserID.String(), accessTokenID.String(), config.Env.Token.AccessTokenDuration)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	refreshTokenID := uuid.New()
	refreshToken, refreshExpiredAt, err := util.GenerateToken(config.Env.Token.RefreshTokenSecret, req.UserID.String(), refreshTokenID.String(), config.Env.Token.RefreshTokenDuration)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	cacheKey := cachekey.NewRefreshTokenCacheKey(req.UserID.String(), req.TokenID.String())
	err = s.cacheRepository.DeleteCache(ctx, cacheKey)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &response.BaseResponse{
		Message: response.MessageOK,
		Data: response.AuthResp{
			AccessToken:           accessToken,
			AccessTokenExpiredAt:  accessExpiredAt.Format(time.RFC3339),
			RefreshToken:          refreshToken,
			RefreshTokenExpiredAt: refreshExpiredAt.Format(time.RFC3339),
		},
	}, nil
}
