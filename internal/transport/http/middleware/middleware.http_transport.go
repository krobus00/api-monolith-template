package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/api-monolith-template/internal/config"
	"github.com/api-monolith-template/internal/constant"
	"github.com/api-monolith-template/internal/model/cachekey"
	"github.com/api-monolith-template/internal/model/contract"
	"github.com/api-monolith-template/internal/model/request"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	authService     contract.AuthService
	cacheRepository contract.CacheRepository
}

func NewController() *Controller {
	return new(Controller)
}

func (c *Controller) WithAuthService(svc contract.AuthService) *Controller {
	c.authService = svc
	return c
}

func (c *Controller) WithCacheRepository(repo contract.CacheRepository) *Controller {
	c.cacheRepository = repo
	return c
}

func (c *Controller) LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now().UTC()
		ctx.Next()
		endTime := time.Now().UTC()
		latencyTime := endTime.Sub(startTime)
		reqMethod := ctx.Request.Method
		reqUri := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()

		logrus.WithFields(logrus.Fields{
			"method":   reqMethod,
			"uri":      reqUri,
			"status":   statusCode,
			"latency":  latencyTime,
			"clientIP": clientIP,
		}).Info("Incoming request")

		ctx.Next()
	}
}

func (c *Controller) CustomPanicHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				stackBuf := make([]byte, 1024)
				stackSize := runtime.Stack(stackBuf, false)
				stackTrace := string(stackBuf[:stackSize])

				logger := logrus.WithFields(logrus.Fields{
					"stackTrace": stackTrace,
				})

				if err, ok := r.(error); ok {
					logger.Error(err)
				} else {
					logger.Error(fmt.Sprintf("panic occurred: %v", r))
				}

				c.JSON(http.StatusInternalServerError, constant.ErrInternalServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}

func (c *Controller) AuthMiddleware(tokenType string) gin.HandlerFunc {
	return func(ginCtx *gin.Context) {

		tokenSecret := config.Env.Token.AccessTokenSecret
		if tokenType == constant.RefreshTokenType {
			tokenSecret = config.Env.Token.RefreshTokenSecret
		}

		cErr := constant.ErrInvalidToken
		authHeader := ginCtx.GetHeader("Authorization")
		if authHeader == "" {
			ginCtx.JSON(cErr.StatusCode, cErr.ToResponse())
			ginCtx.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ginCtx.JSON(cErr.StatusCode, cErr.ToResponse())
			ginCtx.Abort()
			return
		}
		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(tokenSecret), nil
		})

		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				cerr := constant.ErrTokenExpired
				ginCtx.JSON(cerr.StatusCode, cerr.ToResponse())
				ginCtx.Abort()
				return
			}
			ginCtx.JSON(cErr.StatusCode, cErr.ToResponse())
			ginCtx.Abort()
			return
		}

		if !token.Valid {
			ginCtx.JSON(cErr.StatusCode, cErr.ToResponse())
			ginCtx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			ginCtx.JSON(cErr.StatusCode, cErr.ToResponse())
			ginCtx.Abort()
			return
		}

		userID := claims[string(constant.UserID)].(string)
		ginCtx.Set(string(constant.UserID), userID)
		ginCtx.Request = ginCtx.Request.WithContext(context.WithValue(ginCtx.Request.Context(), constant.UserID, userID))

		// check is current user if valid or not
		_, err = c.authService.Info(ginCtx.Request.Context(), &request.AuthInfoReq{
			UserID: uuid.MustParse(userID),
		})
		if err != nil {
			ginCtx.JSON(cErr.StatusCode, cErr.ToResponse())
			ginCtx.Abort()
			return
		}

		tokenID := claims[string(constant.TokenID)].(string)
		ginCtx.Set(string(constant.TokenID), tokenID)
		ginCtx.Request = ginCtx.Request.WithContext(context.WithValue(ginCtx.Request.Context(), constant.TokenID, tokenID))

		if tokenType == constant.RefreshTokenType {
			// check refresh token is valid or not
			cacheKey := cachekey.NewRefreshTokenCacheKey(userID, tokenID)
			existingToken := ""
			err = c.cacheRepository.GetCache(ginCtx.Request.Context(), cacheKey, &existingToken)
			if err != nil {
				ginCtx.JSON(cErr.StatusCode, cErr.ToResponse())
				ginCtx.Abort()
				return
			}
		}

		ginCtx.Next()
	}
}
