package auth

import (
	"github.com/api-monolith-template/internal/constant"
	"github.com/api-monolith-template/internal/model/request"
	"github.com/api-monolith-template/internal/util"
	"github.com/gin-gonic/gin"
)

func (c *Controller) Logout(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := &request.AuthLogoutReq{}

	userID, err := util.GetUserIDFromContext(ctx)
	if err != nil {
		util.HandleError(ginCtx, constant.ErrInvalidToken)
		ginCtx.Abort()
		return
	}

	tokenID, err := util.GetTokenIDFromContext(ctx)
	if err != nil {
		util.HandleError(ginCtx, constant.ErrInvalidToken)
		ginCtx.Abort()
		return
	}

	req.UserID = *userID
	req.TokenID = *tokenID
	resp, err := c.authService.Logout(ctx, req)
	util.HandleResponse(ginCtx, resp, err)
}
