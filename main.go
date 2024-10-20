package main

import (
	"net/http"

	"github.com/rendaman0215/michopa/internal/server"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	mux := server.NewConnectServer()
	http.ListenAndServe(":8080", h2c.NewHandler(mux, &http2.Server{}))
}
