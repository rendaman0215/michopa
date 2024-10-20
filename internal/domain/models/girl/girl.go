package girl

import (
	commonvalobject "github.com/rendaman0215/michopa/internal/domain/models/common_val_object"
	"github.com/rendaman0215/michopa/internal/err"
)

type Girl struct {
	id     commonvalobject.ID
	name   commonvalobject.Fullname
	age    Age
	cup    Cup
	hip    Hip
	height Height
}

type GirlDTO struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Cup    string `json:"cup"`
	Hip    int    `json:"hip"`
	Height int    `json:"height"`
}

// コンストラクタ
func NewGirl(girlData GirlDTO) *Girl {
	return &Girl{}
}

// ゲッター
func (girl *Girl) Id() *commonvalobject.ID {
	return &girl.id
}
func (girl *Girl) Name() *commonvalobject.Fullname {
	return &girl.name
}
func (girl *Girl) Age() *Age {
	return &girl.age
}
func (girl *Girl) Cup() *Cup {
	return &girl.cup
}
func (girl *Girl) Hip() *Hip {
	return &girl.hip
}
func (girl *Girl) Height() *Height {
	return &girl.height
}

// 同一性検証
func (girl *Girl) Equals(obj *Girl) (bool, error) {
	if obj == nil {
		return false, err.NewInternalError("引数でnilが指定されました。")
	}
	result := girl.id.Equals(obj.Id())
	return result, nil
}
