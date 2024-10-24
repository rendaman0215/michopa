package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/rendaman0215/michopa/internal/infrastructure/memory"
	"github.com/rendaman0215/michopa/internal/interface_adapater/controller"
	"github.com/rendaman0215/michopa/internal/interface_adapater/presenter"
	"github.com/rendaman0215/michopa/internal/server"
	"github.com/rendaman0215/michopa/pkg/logger"
	"go.uber.org/fx"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/rs/cors"
)

func main() {
	// Init Log
	logger.InitLogger()

	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			memory.NewMemCache,
			controller.NewGirlController,
			presenter.NewGirlPresentation,
			server.NewConnectServer,
		),
		fx.Invoke(func(server *server.ConnectServer) {
			mux := server.NewServer()
			corsHandler := cors.AllowAll().Handler(h2c.NewHandler(mux, &http2.Server{}))
			http.ListenAndServe(":8080", corsHandler)
		}),
	)

	slog.Info("start program", "author", os.Getenv("AUTHOR"))

	app.Run()
}
