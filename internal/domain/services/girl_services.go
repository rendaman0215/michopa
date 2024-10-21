package services

import (
	"github.com/rendaman0215/michopa/internal/domain/repository"
)

type GirlDomainService struct {
	GirlRepository repository.GirlRepository
}

func (gs *GirlDomainService) IsGirlAlreadyRegisteredByAttributes(firstname string, lastname string) (bool, error) {
	girls, err := gs.GirlRepository.FindByAttributes(firstname, lastname)
	if err != nil {
		return false, err
	}
	return len(girls) > 0, nil
}
