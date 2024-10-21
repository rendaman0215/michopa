package service_interface

import (
	"context"

	girlv1 "buf.build/gen/go/rendaman/personal-schema/protocolbuffers/go/girl/v1"
	"connectrpc.com/connect"
	"github.com/rendaman0215/michopa/internal/domain/models/girl"
)

// presentation
type GirlPresentation interface {
	ListGirl([]*girl.Girl) (girlv1.GirlServiceListResponse, error)
}

// controller
type GirlController interface {
	RegisterGirl(context.Context, *connect.Request[girlv1.GirlServiceRegisterRequest]) (*girl.Girl, error)
}
