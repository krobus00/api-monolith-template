package bootstrap

import (
	"context"
	"fmt"
	"net/http"

	"github.com/api-monolith-template/internal/config"
	"github.com/api-monolith-template/internal/infrastructure"
	cacheRepo "github.com/api-monolith-template/internal/repository/cache"
	userRepo "github.com/api-monolith-template/internal/repository/user"
	authSvc "github.com/api-monolith-template/internal/service/auth"
	httpTransport "github.com/api-monolith-template/internal/transport/http"
	authCtrl "github.com/api-monolith-template/internal/transport/http/auth"
	"github.com/api-monolith-template/internal/util"
	"github.com/sirupsen/logrus"
)

func StartServer() {
	ctx := context.Background()

	// init infra
	infrastructure.InitializeDBConn()
	db, err := infrastructure.DB.DB()
	util.ContinueOrFatal(err)

	err = db.Ping()
	util.ContinueOrFatal(err)

	rdb := infrastructure.NewRedisClient()
	_, err = rdb.Ping(ctx).Result()
	util.ContinueOrFatal(err)

	r := infrastructure.NewGinEngine()

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
		WithUserRepository(userRepository)

	// init controller
	authController := authCtrl.
		NewController().
		WithAuthService(authService)

	// init http transport
	httpTransport.
		NewTransport().
		WithGinEngine(r).
		WithAuthController(authController).
		InitRoute()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Env.Server.Port),
		Handler: r.Handler(),
	}
	// start http server
	go func() {
		logrus.Info(fmt.Sprintf("running at http://0.0.0.0:%s", config.Env.Server.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			util.ContinueOrFatal(err)
		}
	}()

	wait := gracefulShutdown(ctx, config.Env.GracefulShutdownTimeout, map[string]operation{
		"database connection": func(ctx context.Context) error {
			infrastructure.StopTickerCh <- true
			return db.Close()
		},
		"redis connection": func(ctx context.Context) error {
			return rdb.Close()
		},
		"gin server": func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	<-wait
}
