package auth

import (
	"github.com/api-monolith-template/internal/model/request"
	"github.com/api-monolith-template/internal/util"
	"github.com/gin-gonic/gin"
)

func (c *Controller) Login(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := &request.LoginReq{}

	err := ginCtx.ShouldBindJSON(&req)
	if err != nil {
		util.HandleError(ginCtx, err)
		ginCtx.Abort()
		return
	}

	resp, err := c.authService.Login(ctx, req)
	if err != nil {
		util.HandleError(ginCtx, err)
		ginCtx.Abort()
		return
	}

	ginCtx.JSON(resp.StatusCode, resp)
}
