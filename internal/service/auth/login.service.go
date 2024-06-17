package auth

import (
	"context"
	"errors"
	"time"

	"github.com/api-monolith-template/internal/config"
	"github.com/api-monolith-template/internal/constant"
	"github.com/api-monolith-template/internal/model/request"
	"github.com/api-monolith-template/internal/model/response"
	"github.com/api-monolith-template/internal/util"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (s *Service) Login(ctx context.Context, req *request.LoginReq) (*response.BaseResponse, error) {
	logger := util.NewDefaultLogger(ctx).WithFields(logrus.Fields{
		"req": req,
	})

	user, err := s.userRepository.FindByIdentifier(ctx, req.Identifier)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constant.ErrUserNotFound
	}
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	isPasswordMatch, err := util.ComparePassword(user.Password, req.Password)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if !isPasswordMatch {
		return nil, constant.ErrPasswordNotMatch
	}

	accessTokenID := uuid.New()
	accessToken, accessExpiredAt, err := util.GenerateToken(config.Env.Token.AccessTokenSecret, user.ID.String(), accessTokenID.String(), config.Env.Token.AccessTokenDuration)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	refreshTokenID := uuid.New()
	refreshToken, refreshExpiredAt, err := util.GenerateToken(config.Env.Token.RefreshTokenSecret, user.ID.String(), refreshTokenID.String(), config.Env.Token.RefreshTokenDuration)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	// TODO: store refresh token

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
