package repository

import (
	"context"

	"github.com/rendaman0215/michopa/internal/domain/models/girl"
)

// GirlRepository is a repository interface
type GirlRepository interface {
	// 女の子一覧を取得する
	List(ctx context.Context) ([]*girl.Girl, error)

	// 女の子を登録する
	Register(ctx context.Context, girl *girl.Girl) error
}
