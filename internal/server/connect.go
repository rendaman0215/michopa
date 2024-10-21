package server

import (
	"net/http"

	"buf.build/gen/go/rendaman/personal-schema/connectrpc/go/girl/v1/girlv1connect"
	"connectrpc.com/connect"
	"github.com/rendaman0215/michopa/internal/domain/repository"
	"github.com/rendaman0215/michopa/internal/domain/services"
	"github.com/rendaman0215/michopa/internal/server/middleware"
	"github.com/rendaman0215/michopa/internal/service"
	service_interface "github.com/rendaman0215/michopa/internal/service/interface"
)

type ConnectServer struct {
	*service.GirlService
}

func NewConnectServer(girlRepo repository.GirlRepository, girlPresen service_interface.GirlPresentation, girlCon service_interface.GirlController) *ConnectServer {
	girlDomainService := services.GirlDomainService{
		GirlRepository: girlRepo,
	}
	gs := service.NewGirlService(girlRepo, girlPresen, girlCon, girlDomainService)
	return &ConnectServer{
		GirlService: gs,
	}
}

func (cs *ConnectServer) NewServer() *http.ServeMux {
	logInterceptor := middleware.NewLoggerInterceptor()
	interceptors := connect.WithInterceptors(logInterceptor)

	mux := http.NewServeMux()
	path, handler := girlv1connect.NewGirlServiceHandler(cs.GirlService, interceptors)
	mux.Handle(path, handler)

	return mux
}
