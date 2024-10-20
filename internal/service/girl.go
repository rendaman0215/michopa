package service

import (
	"context"

	girlv1 "buf.build/gen/go/rendaman/personal-schema/protocolbuffers/go/girl/v1"
	"connectrpc.com/connect"
)

type GirlService struct{}

func NewGirlService() *GirlService {
	return &GirlService{}
}

func (gs *GirlService) List(ctx context.Context, req *connect.Request[girlv1.GirlServiceListRequest]) (*connect.Response[girlv1.GirlServiceListResponse], error) {
	return connect.NewResponse(&girlv1.GirlServiceListResponse{
		Girl: []*girlv1.Girl{},
	}), nil
}

func (gs *GirlService) Register(ctx context.Context, req *connect.Request[girlv1.GirlServiceRegisterRequest]) (*connect.Response[girlv1.GirlServiceRegisterResponse], error) {
	return connect.NewResponse(&girlv1.GirlServiceRegisterResponse{}), nil
}
