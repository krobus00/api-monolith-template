package contract

import (
	"context"

	"github.com/api-monolith-template/internal/model/request"
	"github.com/api-monolith-template/internal/model/response"
)

type AuthService interface {
	Register(ctx context.Context, req *request.RegisterReq) (*response.BaseResponse, error)
}
