package bootstrap

import (
	"context"
	"fmt"
	"net/http"

	"github.com/api-monolith-template/internal/config"
	"github.com/api-monolith-template/internal/infrastructure"
	authSvc "github.com/api-monolith-template/internal/service/auth"
	httpTransport "github.com/api-monolith-template/internal/transport/http"
	authCtrl "github.com/api-monolith-template/internal/transport/http/auth"
	"github.com/api-monolith-template/internal/util"
	"github.com/sirupsen/logrus"
)

func StartServer() {
	// init infra
	infrastructure.InitializeDBConn()
	db, err := infrastructure.DB.DB()
	util.ContinueOrFatal(err)

	err = db.Ping()
	util.ContinueOrFatal(err)

	r := infrastructure.NewGinEngine()

	// init service
	authService := authSvc.NewService()

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

	wait := gracefulShutdown(context.Background(), config.Env.GracefulShutdownTimeout, map[string]operation{
		"database connection": func(ctx context.Context) error {
			infrastructure.StopTickerCh <- true
			return db.Close()
		},
		"gin server": func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	<-wait
}
