package infrastructure

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/api-monolith-template/internal/config"
	"github.com/api-monolith-template/internal/constant"
	"github.com/api-monolith-template/internal/model/response"
	"github.com/api-monolith-template/internal/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewGinEngine() *gin.Engine {
	r := gin.New()

	if config.Env.Env == constant.ProductionEnvironment {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowOrigins:           nil,
		AllowOriginFunc:        nil,
		AllowMethods:           []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:           []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials:       true,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		AllowFiles:             true,
	}))

	// register custom validation
	util.AddValidation(DB)

	// handle health check
	internalGroup := r.Group("/_internal")
	internalGroup.GET("/healthz", func(c *gin.Context) {
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		status := "healthy"

		serviceStatuses := make([]response.GetHealthCheckServiceStatusResp, 0)

		// check all infrastructure health check
		for serviceName, healthCheckFn := range MapHealthCheck {
			err := healthCheckFn(c.Request.Context())
			if err != nil {
				status = "unhealthy"
			}
			serviceStatuses = append(serviceStatuses, response.GetHealthCheckServiceStatusResp{
				Name: serviceName,
				IsUp: err == nil,
			})
		}

		healthInfo := response.GetHealthCheckResp{
			Status:       status,
			Environtment: config.Env.Env,
			Version:      fmt.Sprintf("%s@%s", config.ServiceName, config.ServiceVersion),
			GoVersion:    runtime.Version(),
			GoRoutine:    runtime.NumGoroutine(),
			Memory: response.GetHealthCheckMemoryResp{
				Alloc:      memStats.Alloc,
				TotalAlloc: memStats.TotalAlloc,
				Sys:        memStats.Sys,
				HeapAlloc:  memStats.HeapAlloc,
				HeapSys:    memStats.HeapSys,
			},
			ServiceStatuses: serviceStatuses,
		}

		resp := response.BaseResponse{
			StatusCode: http.StatusOK,
			Message:    "Health Check",
			Data:       healthInfo,
		}
		c.JSON(resp.StatusCode, resp)
	})

	return r
}
