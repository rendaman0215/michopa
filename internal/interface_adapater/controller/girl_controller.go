package controller

import (
	"context"

	girlv1 "buf.build/gen/go/rendaman/personal-schema/protocolbuffers/go/girl/v1"
	"connectrpc.com/connect"
	"github.com/rendaman0215/michopa/internal/domain/models/girl"
	"github.com/rendaman0215/michopa/internal/interface_adapater/dto"
	service_interface "github.com/rendaman0215/michopa/internal/service/interface"
)

type GirlController struct{}

func NewGirlController() service_interface.GirlController {
	return &GirlController{}
}

func (gc *GirlController) RegisterGirl(ctx context.Context, req *connect.Request[girlv1.GirlServiceRegisterRequest]) (*girl.Girl, error) {
	newGirl, err := girl.NewGirl(ctx, dto.GirlDTO{
		ID:        req.Msg.Girl.Id,
		Firstname: req.Msg.Girl.Name.Firstname,
		Lastname:  req.Msg.Girl.Name.Lastname,
		Age:       int(req.Msg.Girl.Age),
		Cup:       req.Msg.Girl.Cup,
		Hip:       int(req.Msg.Girl.Hip),
		Height:    int(req.Msg.Girl.Height),
	})
	if err != nil {
		return nil, err
	}
	return newGirl, nil
}