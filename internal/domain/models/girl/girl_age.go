package girl

import "github.com/rendaman0215/michopa/internal/err"

type Age struct {
	value int
}

func NewAge(value int) (*Age, error) {
	// 年齢が18歳以上であるかどうか
	if value < 18 {
		return nil, err.NewInternalError("年齢は18歳以上でなければなりません。")
	}

	return &Age{value: value}, nil
}

func (ins *Age) Value() int {
	return ins.value
}

func (ins *Age) Equals(value *Age) bool {
	if ins == value {
		return true
	}
	return ins.value == value.Value()
}
