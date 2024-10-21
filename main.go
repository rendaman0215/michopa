package main

import (
	"net/http"

	"github.com/rendaman0215/michopa/internal/infrastructure/memory"
	"github.com/rendaman0215/michopa/internal/interface_adapater/controller"
	"github.com/rendaman0215/michopa/internal/interface_adapater/presenter"
	"github.com/rendaman0215/michopa/internal/server"
	"github.com/rendaman0215/michopa/pkg/logger"
	"go.uber.org/fx"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	// Init Log
	logger.InitLogger()

	app := fx.New(
		fx.Provide(
			memory.NewMemCache,
			controller.NewGirlController,
			presenter.NewGirlPresentation,
			server.NewConnectServer,
		),
		fx.Invoke(func(server *server.ConnectServer) {
			http.ListenAndServe(":8080", h2c.NewHandler(server.NewServer(), &http2.Server{}))
		}),
	)
	app.Run()
}
