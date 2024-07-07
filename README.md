# Monolith Service

## Response
```json
{
	"message": "validation error", // string
	"data": {
		"id": "6759e56a-fa7c-49ee-9854-f32ab38083ae",
		"username": "username17",
		"email": "email17@gmail.com",
		"level": "USER",
		"createdAt": "2024-06-16T10:46:25Z",
		"updatedAt": "2024-06-16T10:46:25Z"
	}, // object | array of object | null
	"validationErrors": [
		{
			"field": "username",
			"tag": "required",
			"message": "Key: 'RegisterReq.Username' Error:Field validation for 'username' failed on the 'required' tag"
		}
	] // array of object | null
}
```

## API Contract

### Register
#### Request

**Method:** `POST`

**URL:** `${{HOST}}/v1/auth/register`

**Headers:**
- `Content-Type: application/json`

**Body:**
```json
{
	"username":"username",
	"email": "email@gmail.com",
	"password": "strongpassword"
}
```

**Example cURL Command:**

```bash
curl --request POST \
  --url ${{HOST}}/v1/auth/register \
  --header 'Content-Type: application/json' \
  --data '{
	"username":"username",
	"email": "email@gmail.com",
	"password": "strongpassword"
}'
```

**Example Response:**
```json
{
	"message": "validation error",
	"data": null,
	"validationErrors": [
		{
			"field": "username",
			"tag": "unique_db",
			"message": "username already taken"
		},
		{
			"field": "email",
			"tag": "unique_db",
			"message": "email already taken"
		}
	]
}
```

```json
{
	"message": "ok",
	"data": null,
	"validationErrors": null
}
```

### Login
#### Request

**Method:** `POST`

**URL:** `${{HOST}}/v1/auth/login`

**Headers:**
- `Content-Type: application/json`

**Body:**
```json
{
	"identifier":"username",
	"password": "strongpassword"
}
```

**Example cURL Command:**

```bash
curl --request POST \
  --url ${{HOST}}/v1/auth/login \
  --header 'Content-Type: application/json' \
  --data '{
	"identifier":"username",
	"password": "strongpassword"
}'
```

**Example Response:**
```json
{
	"message": "user not found",
	"data": null,
	"validationErrors": null
}
```

```json
{
	"message": "ok",
	"data": {
		"accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI2NzU5ZTU2YS1mYTdjLTQ5ZWUtOTg1NC1mMzJhYjM4MDgzYWUiLCJleHAiOjE3MTg2MTk3ODQsImlhdCI6MTcxODYxNjE4NCwianRpIjoiYzViZjAyZTctNmNkMy00MjZiLThiMjctYzk2MTUyZjc2NmU4In0.aaUAM7Hl6Z-H8kzdnrLedVmmVJEuglxes7xQYHt1HKI",
		"accessTokenExpiredAt": "2024-06-17T09:23:04Z",
		"refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI2NzU5ZTU2YS1mYTdjLTQ5ZWUtOTg1NC1mMzJhYjM4MDgzYWUiLCJleHAiOjE3MTg3MjQxODQsImlhdCI6MTcxODYxNjE4NCwianRpIjoiNzllNTVkZDgtMzczMS00OWU2LThjZDItNzMxNDI0MzYzZjZjIn0.KQOZGZxz8-8JiJv68Xpdj-7z1Dp6dLe0a4IC0nZ5WcA",
		"refreshTokenExpiredAt": "2024-06-17T09:23:04Z"
	},
	"validationErrors": null
}
```

### Refresh Token
#### Request

**Method:** `POST`

**URL:** `${{HOST}}/v1/auth/refresh`

**Headers:**
- `Content-Type: application/json`
- `Authorization: Bearer <REFRESH_TOKEN>`

**Example cURL Command:**

```bash
curl --request POST \
  --url ${{HOST}}/v1/auth/refresh \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer <REFRESH_TOKEN>'
```

**Example Response:**
```json
{
	"message": "invalid token",
	"data": null,
	"validationErrors": null
}
```

```json
{
	"message": "ok",
	"data": {
		"accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI2NzU5ZTU2YS1mYTdjLTQ5ZWUtOTg1NC1mMzJhYjM4MDgzYWUiLCJleHAiOjE3MTg2MTk3ODQsImlhdCI6MTcxODYxNjE4NCwianRpIjoiYzViZjAyZTctNmNkMy00MjZiLThiMjctYzk2MTUyZjc2NmU4In0.aaUAM7Hl6Z-H8kzdnrLedVmmVJEuglxes7xQYHt1HKI",
		"accessTokenExpiredAt": "2024-06-17T09:23:04Z",
		"refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI2NzU5ZTU2YS1mYTdjLTQ5ZWUtOTg1NC1mMzJhYjM4MDgzYWUiLCJleHAiOjE3MTg3MjQxODQsImlhdCI6MTcxODYxNjE4NCwianRpIjoiNzllNTVkZDgtMzczMS00OWU2LThjZDItNzMxNDI0MzYzZjZjIn0.KQOZGZxz8-8JiJv68Xpdj-7z1Dp6dLe0a4IC0nZ5WcA",
		"refreshTokenExpiredAt": "2024-06-17T09:23:04Z"
	},
	"validationErrors": null
}
```

### Logout
#### Request

**Method:** `POST`

**URL:** `${{HOST}}/v1/auth/logout`

**Headers:**
- `Content-Type: application/json`
- `Authorization: Bearer <ACCESS_TOKEN>`

**Example cURL Command:**

```bash
curl --request POST \
  --url ${{HOST}}/v1/auth/logout \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer <ACCESS_TOKEN>'
```

**Example Response:**
```json
{
	"message": "invalid token",
	"data": null,
	"validationErrors": null
}
```

```json
{
	"message": "ok",
	"data": null,
	"validationErrors": null
}
```

### User Info
#### Request

**Method:** `GET`

**URL:** `${{HOST}}/v1/auth/info`

**Headers:**
- `Content-Type: application/json`
- `Authorization: Bearer <ACCESS_TOKEN>`

**Example cURL Command:**

```bash
curl --request GET \
  --url ${{HOST}}/v1/auth/info \
  --header 'Authorization: Bearer <ACCESS_TOKEN>'
```

**Example Response:**
```json
{
	"message": "ok",
	"data": {
		"id": "6759e56a-fa7c-49ee-9854-f32ab38083ae",
		"username": "username",
		"email": "email@gmail.com",
		"level": "USER",
		"createdAt": "2024-06-16T10:46:25Z",
		"updatedAt": "2024-06-16T10:46:25Z"
	},
	"validationErrors": null
}
```

## Create new domain

### Define contract

Define your contract domain (repository and service) on /internal/model/contract/<DOMAIN>.contract.go, example:

```go
package contract

import (
	"context"

	"github.com/api-monolith-template/internal/model/entity"
    "github.com/api-monolith-template/internal/model/request"
	"github.com/api-monolith-template/internal/model/response"
	"github.com/google/uuid"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	FindByIdentifier(ctx context.Context, identifier string) (*entity.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	Upsert(ctx context.Context, user *entity.User) error
}

type AuthService interface {
	Register(ctx context.Context, req *request.RegisterReq) (*response.BaseResponse, error)
	Login(ctx context.Context, req *request.LoginReq) (*response.BaseResponse, error)
	RefreshToken(ctx context.Context, req *request.AuthRefreshReq) (*response.BaseResponse, error)
	Info(ctx context.Context, req *request.AuthInfoReq) (*response.BaseResponse, error)
	Logout(ctx context.Context, req *request.AuthLogoutReq) (*response.BaseResponse, error)
}

```

### Create your repository

Create your base repository on /internal/repository/<DOMAIN>/<DOMAIN>.repository.go, example:

```go
package user

import (
	"github.com/api-monolith-template/internal/model/contract"
	"gorm.io/gorm"
)

type Repository struct {
	db        *gorm.DB
	cacheRepo contract.CacheRepository
}

func NewRepository() *Repository {
	return new(Repository)
}

func (r *Repository) WithGormDB(db *gorm.DB) *Repository {
	r.db = db
	return r
}

func (r *Repository) WithCacheRepository(repo contract.CacheRepository) *Repository {
	r.cacheRepo = repo
	return r
}
```

### Create your service

Create your base service on /internal/service/<DOMAIN>/<DOMAIN>.service.go, example:

```go
package auth

import "github.com/api-monolith-template/internal/model/contract"

type Service struct {
	userRepository  contract.UserRepository
	cacheRepository contract.CacheRepository
}

func NewService() *Service {
	return new(Service)
}

func (s *Service) WithUserRepository(repo contract.UserRepository) *Service {
	s.userRepository = repo
	return s
}

func (s *Service) WithCacheRepository(repo contract.CacheRepository) *Service {
	s.cacheRepository = repo
	return s
}
```

### Create your transport layer

Create base http transport layer on /internal/transport/http/<DOMAIN>/<DOMAIN>.http_transport.go, example:

```go
package auth

import "github.com/api-monolith-template/internal/model/contract"

type Controller struct {
	authService contract.AuthService
}

func NewController() *Controller {
	return new(Controller)
}

func (c *Controller) WithAuthService(svc contract.AuthService) *Controller {
	c.authService = svc
	return c
}
```

inject all your transport domain on /internal/transport/http/http.transport.go

```go
package http

import (
	"github.com/api-monolith-template/internal/transport/http/auth"
	"github.com/api-monolith-template/internal/transport/http/middleware"
	"github.com/gin-gonic/gin"
)

type Transport struct {
	router *gin.Engine

	authController       *auth.Controller
	middlewareController *middleware.Controller
}

func NewTransport() *Transport {
	return new(Transport)
}

func (t *Transport) WithGinEngine(r *gin.Engine) *Transport {
	t.router = r
	return t
}

func (t *Transport) WithAuthController(c *auth.Controller) *Transport {
	t.authController = c
	return t
}

func (t *Transport) WithMiddlewareController(c *middleware.Controller) *Transport {
	t.middlewareController = c
	return t
}

```

after create http transport layer, create http route on /internal/transport/http/route.transport.go

```go
authGroup := v1Group.Group("/auth")
authGroup.POST("/register", t.authController.Register)
authGroup.POST("/login", t.authController.Login)
authRefreshToken := authGroup.Group("/refresh", t.middlewareController.AuthMiddleware(constant.RefreshTokenType))
authRefreshToken.POST("/", t.authController.RefreshToken)

authProtected := authGroup.Use(t.middlewareController.AuthMiddleware(constant.AccessTokenType))
authProtected.GET("/info", t.authController.Info)
```

### Inject all new depedency

after create a domain for each layer, now init new domain and inject to all layer

```go
// init repository
cacheRepository := cacheRepo.
    NewRepository().
    WithRedisDB(rdb)
userRepository := userRepo.
    NewRepository().
    WithGormDB(infrastructure.DB).
    WithCacheRepository(cacheRepository)

// init service
authService := authSvc.
    NewService().
    WithUserRepository(userRepository).
    WithCacheRepository(cacheRepository)

// init controller
middlewareController := middlewareCtrl.
    NewController().
    WithAuthService(authService).
    WithCacheRepository(cacheRepository)
authController := authCtrl.
    NewController().
    WithAuthService(authService)

// init http transport
httpTransport.
    NewTransport().
    WithGinEngine(r).
    WithMiddlewareController(middlewareController).
    WithAuthController(authController).
    InitRoute()
```
