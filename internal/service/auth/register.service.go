package auth

import (
	"context"

	"github.com/api-monolith-template/internal/model/request"
	"github.com/api-monolith-template/internal/model/response"
)

func (s *Service) Register(ctx context.Context, req *request.RegisterReq) (*response.BaseResponse, error) {
	return response.NewResponseOK(), nil
}
