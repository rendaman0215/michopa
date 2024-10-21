package memory

import (
	"context"

	"github.com/rendaman0215/michopa/internal/domain/models/girl"
	"github.com/rendaman0215/michopa/internal/domain/repository"
)

type MemRepo struct {
	girls []*girl.Girl
}

func NewMemCache() repository.GirlRepository {
	return &MemRepo{}
}

func (mr *MemRepo) Register(ctx context.Context, girl *girl.Girl) error {
	mr.girls = append(mr.girls, girl)
	return nil
}

func (mr *MemRepo) List(ctx context.Context) ([]*girl.Girl, error) {
	return mr.girls, nil
}

// FindByAttributes は属性で検索する
func (mr *MemRepo) FindByAttributes(firstname string, lastname string) ([]*girl.Girl, error) {
	var result []*girl.Girl
	for _, g := range mr.girls {
		if g.Name().Firstname() == firstname && g.Name().Lastname() == lastname {
			result = append(result, g)
		}
	}
	return result, nil
}
