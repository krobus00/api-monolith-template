package bootstrap

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/api-monolith-template/internal/config"
	"github.com/api-monolith-template/internal/infrastructure"
	httpTransport "github.com/api-monolith-template/internal/transport/http"
	"github.com/sirupsen/logrus"
)

func StartServer() {
	r := infrastructure.NewGinEngine()

	httpTransport.
		NewTransport().
		WithGinEngine(r).
		InitRoute()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Env.Server.Port),
		Handler: r.Handler(),
	}
	// start http server
	go func() {
		logrus.Info(fmt.Sprintf("running at http://0.0.0.0:%s", config.Env.Server.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatal(err)
		}
	}()

	wait := gracefulShutdown(context.Background(), 15*time.Second, map[string]operation{
		"gin server": func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	<-wait
}
