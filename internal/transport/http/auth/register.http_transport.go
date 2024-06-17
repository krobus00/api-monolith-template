package auth

import (
	"github.com/api-monolith-template/internal/model/request"
	"github.com/api-monolith-template/internal/util"
	"github.com/gin-gonic/gin"
)

func (c *Controller) Register(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := &request.RegisterReq{}

	err := ginCtx.ShouldBindJSON(&req)
	if err != nil {
		util.HandleError(ginCtx, err)
		ginCtx.Abort()
		return
	}

	resp, err := c.authService.Register(ctx, req)
	util.HandleResponse(ginCtx, resp, err)
}
