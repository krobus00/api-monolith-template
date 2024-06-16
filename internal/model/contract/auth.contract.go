package contract

import (
	"context"

	"github.com/api-monolith-template/internal/model/request"
	"github.com/api-monolith-template/internal/model/response"
)

type AuthService interface {
	Register(ctx context.Context, req *request.RegisterReq) (*response.BaseResponse, error)
	Login(ctx context.Context, req *request.LoginReq) (*response.BaseResponse, error)
	Info(ctx context.Context, req *request.AuthInfoReq) (*response.BaseResponse, error)
}
