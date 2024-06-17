package auth

import (
	"context"

	"github.com/api-monolith-template/internal/config"
	"github.com/api-monolith-template/internal/constant"
	"github.com/api-monolith-template/internal/model/entity"
	"github.com/api-monolith-template/internal/model/request"
	"github.com/api-monolith-template/internal/model/response"
	"github.com/api-monolith-template/internal/util"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (s *Service) Register(ctx context.Context, req *request.RegisterReq) (*response.BaseResponse, error) {
	logger := util.NewDefaultLogger(ctx).WithFields(logrus.Fields{
		"req": util.Dump(req),
	})

	hashPassword, err := util.HashPassword(req.Password, []byte(config.Env.Token.PasswordSalt))
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	newUser := &entity.User{
		ID:       uuid.New(),
		Username: req.Username,
		Email:    req.Email,
		Password: hashPassword,
		Level:    constant.LevelUser,
	}

	err = s.userRepository.Upsert(ctx, newUser)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return response.NewResponseOK(), nil
}
