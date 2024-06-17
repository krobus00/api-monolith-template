package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/api-monolith-template/internal/constant"
	"github.com/api-monolith-template/internal/model/request"
	"github.com/api-monolith-template/internal/model/response"
	"github.com/api-monolith-template/internal/util"
	"gorm.io/gorm"
)

func (s *Service) Info(ctx context.Context, req *request.AuthInfoReq) (*response.BaseResponse, error) {
	logger := util.NewDefaultLogger(ctx)

	user, err := s.userRepository.FindByID(ctx, req.UserID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Error(err)
		return nil, err
	}
	if user == nil {
		return nil, constant.ErrUserNotFound
	}

	return &response.BaseResponse{
		StatusCode: http.StatusOK,
		Message:    response.MessageOK,
		Data: response.AuthInfoResp{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Level:     user.Level,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		},
	}, nil
}
