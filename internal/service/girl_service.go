package service

import (
	"context"
	"fmt"

	girlv1 "buf.build/gen/go/rendaman/personal-schema/protocolbuffers/go/girl/v1"
	"connectrpc.com/connect"
	"github.com/rendaman0215/michopa/internal/domain/repository"
	"github.com/rendaman0215/michopa/internal/domain/services"
	service_interface "github.com/rendaman0215/michopa/internal/service/interface"
)

type GirlService struct {
	repo          repository.GirlRepository
	presenter     service_interface.GirlPresentation
	controller    service_interface.GirlController
	domainService services.GirlDomainService
}

func NewGirlService(girlRepo repository.GirlRepository, girlPresen service_interface.GirlPresentation, girlCon service_interface.GirlController, girlDomainService services.GirlDomainService) *GirlService {
	return &GirlService{
		repo:          girlRepo,
		presenter:     girlPresen,
		controller:    girlCon,
		domainService: girlDomainService,
	}
}

func (gs *GirlService) List(ctx context.Context, req *connect.Request[girlv1.GirlServiceListRequest]) (*connect.Response[girlv1.GirlServiceListResponse], error) {
	girls, err := gs.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	res, err := gs.presenter.ListGirl(girls)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&res), nil
}

func (gs *GirlService) Register(ctx context.Context, req *connect.Request[girlv1.GirlServiceRegisterRequest]) (*connect.Response[girlv1.GirlServiceRegisterResponse], error) {
	newGirl, err := gs.controller.RegisterGirl(ctx, req)
	if err != nil {
		return nil, err
	}

	isRegistered, err := gs.domainService.IsGirlAlreadyRegisteredByAttributes(newGirl.Name().Firstname(), newGirl.Name().Lastname())
	if err != nil {
		return nil, err
	}

	if isRegistered {
		return nil, connect.NewError(connect.CodeAlreadyExists, fmt.Errorf("This girl is already registered"))
	}

	err = gs.repo.Register(ctx, newGirl)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&girlv1.GirlServiceRegisterResponse{}), nil
}
