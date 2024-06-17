package auth

import (
	"context"

	"github.com/api-monolith-template/internal/config"
	"github.com/api-monolith-template/internal/model/cachekey"
	"github.com/api-monolith-template/internal/model/request"
	"github.com/api-monolith-template/internal/model/response"
	"github.com/api-monolith-template/internal/repository/cache"
	"github.com/api-monolith-template/internal/util"
	"github.com/sirupsen/logrus"
)

func (s *Service) Logout(ctx context.Context, req *request.AuthLogoutReq) (*response.BaseResponse, error) {
	logger := util.NewDefaultLogger(ctx).WithFields(logrus.Fields{
		"req": req,
	})

	// delete old refresh token
	cacheKey := cachekey.NewRefreshTokenCacheKey(req.UserID.String(), req.TokenID.String())
	err := s.cacheRepository.DeleteCache(ctx, cacheKey)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	// mark old access token as blacklist
	cacheKey = cachekey.NewAccessTokenBlacklistCacheKey(req.UserID.String(), req.TokenID.String())
	err = s.cacheRepository.SetCache(ctx, cacheKey, req.TokenID.String(), cache.WithCustomExpiredDuration(config.Env.Token.AccessTokenDuration))
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return response.NewResponseOK(), nil
}
