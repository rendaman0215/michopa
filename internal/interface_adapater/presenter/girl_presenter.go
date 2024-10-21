package presenter

import (
	girlv1 "buf.build/gen/go/rendaman/personal-schema/protocolbuffers/go/girl/v1"
	"github.com/rendaman0215/michopa/internal/domain/models/girl"
	service_interface "github.com/rendaman0215/michopa/internal/service/interface"
)

type GirlPresentation struct{}

func NewGirlPresentation() service_interface.GirlPresentation {
	return &GirlPresentation{}
}

func (gp *GirlPresentation) ListGirl(girls []*girl.Girl) (girlv1.GirlServiceListResponse, error) {
	var res []*girlv1.Girl
	for _, targetGirl := range girls {
		res = append(res, &girlv1.Girl{
			Id: targetGirl.Id().Value(),
			Name: &girlv1.Name{
				Firstname: targetGirl.Name().Firstname(),
				Lastname:  targetGirl.Name().Lastname(),
			},
			Age:    int32(targetGirl.Age().Value()),
			Cup:    targetGirl.Cup().Value(),
			Hip:    int32(targetGirl.Hip().Value()),
			Height: int32(targetGirl.Height().Value()),
		})
	}

	return girlv1.GirlServiceListResponse{
		Girl: res,
	}, nil
}
