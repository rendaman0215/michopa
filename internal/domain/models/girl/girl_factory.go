package girl

import (
	"context"

	"github.com/google/uuid"
	commonvalobject "github.com/rendaman0215/michopa/internal/domain/models/common_val_object"
	"github.com/rendaman0215/michopa/internal/interface_adapater/dto"
	"golang.org/x/sync/errgroup"
)

// コンストラクタ
func NewGirl(ctx context.Context, girlData dto.GirlDTO) (*Girl, error) {
	var (
		id     *commonvalobject.ID
		name   *commonvalobject.Fullname
		age    *Age
		cup    *Cup
		hip    *Hip
		height *Height
	)

	// エラグループを作成して並行処理を行う
	g, ctx := errgroup.WithContext(ctx)

	// idをセット
	g.Go(func() error {
		var err error
		uid, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		id, err = commonvalobject.NewId(uid.String())
		if err != nil {
			return err
		}
		return nil
	})

	// nameをセット
	g.Go(func() error {
		var err error
		name, err = commonvalobject.NewFullname(girlData.Firstname, girlData.Lastname)
		if err != nil {
			return err
		}
		return nil
	})

	// ageをセット
	g.Go(func() error {
		var err error
		age, err = NewAge(girlData.Age)
		if err != nil {
			return err
		}
		return nil
	})

	// cupをセット
	g.Go(func() error {
		var err error
		cup, err = NewCup(girlData.Cup)
		if err != nil {
			return err
		}
		return nil
	})

	// hipをセット
	g.Go(func() error {
		var err error
		hip, err = NewHip(girlData.Hip)
		if err != nil {
			return err
		}
		return nil
	})

	// heightをセット
	g.Go(func() error {
		var err error
		height, err = NewHeight(girlData.Height)
		if err != nil {
			return err
		}
		return nil
	})

	// 並行処理が完了するのを待つ
	if err := g.Wait(); err != nil {
		return nil, err
	}

	// 全てのフィールドが正しく設定されたらGirlインスタンスを返す
	return &Girl{
		id:     *id,
		name:   *name,
		age:    *age,
		cup:    *cup,
		hip:    *hip,
		height: *height,
	}, nil
}
