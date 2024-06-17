package auth

import (
	"github.com/api-monolith-template/internal/constant"
	"github.com/api-monolith-template/internal/model/request"
	"github.com/api-monolith-template/internal/util"
	"github.com/gin-gonic/gin"
)

func (c *Controller) Info(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := &request.AuthInfoReq{}

	userID, err := util.GetUserIDFromContext(ctx)
	if err != nil {
		util.HandleError(ginCtx, constant.ErrInvalidToken)
		ginCtx.Abort()
		return
	}

	req.UserID = *userID
	resp, err := c.authService.Info(ctx, req)
	util.HandleResponse(ginCtx, resp, err)
}
