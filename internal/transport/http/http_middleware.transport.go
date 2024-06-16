package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/api-monolith-template/internal/constant"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

func loggingMiddleware() gin.HandlerFunc {
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

func customPanicHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					logrus.Error(err)
				}
				c.JSON(http.StatusInternalServerError, constant.ErrInternalServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		cErr := constant.ErrInvalidToken
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(cErr.StatusCode, cErr.ToResponse())
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(cErr.StatusCode, cErr.ToResponse())
			c.Abort()
			return
		}
		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				cerr := constant.ErrTokenExpired
				c.JSON(cerr.StatusCode, cerr.ToResponse())
				c.Abort()
				return
			}
			c.JSON(cErr.StatusCode, cErr.ToResponse())
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(cErr.StatusCode, cErr.ToResponse())
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims[string(constant.UserID)].(string)
			c.Set(string(constant.UserID), userID)
			c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), constant.UserID, userID))
			c.Next()
		} else {
			c.JSON(cErr.StatusCode, cErr.ToResponse())
			c.Abort()
			return
		}
	}
}
