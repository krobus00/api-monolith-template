package auth

import (
	"github.com/api-monolith-template/internal/constant"
	"github.com/api-monolith-template/internal/model/request"
	"github.com/api-monolith-template/internal/util"
	"github.com/gin-gonic/gin"
)

func (c *Controller) RefreshToken(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := &request.AuthRefreshReq{}

	userID, err := util.GetUserIDFromContext(ctx)
	if err != nil {
		util.HandleError(ginCtx, constant.ErrInvalidToken)
		ginCtx.Abort()
		return
	}

	req.UserID = *userID
	resp, err := c.authService.RefreshToken(ctx, req)
	util.HandleResponse(ginCtx, resp, err)
}
