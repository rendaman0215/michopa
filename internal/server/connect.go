package server

import (
	"net/http"

	"buf.build/gen/go/rendaman/personal-schema/connectrpc/go/girl/v1/girlv1connect"
	"github.com/rendaman0215/michopa/internal/service"
)

type ConnectServer struct{}

func NewConnectServer() *http.ServeMux {
	girlService := service.NewGirlService()
	mux := http.NewServeMux()
	path, handler := girlv1connect.NewGirlServiceHandler(girlService)
	mux.Handle(path, handler)

	return mux
}
