package girl

import "github.com/rendaman0215/michopa/internal/err"

type Height struct {
	value int
}

func NewHeight(value int) (*Height, error) {
	// 身長が130cm以上、200cm以下であるかどうか
	if value < 130 || value > 200 {
		return nil, err.NewInternalError("身長は130cm以上、200cm以下でなければなりません。")
	}

	return &Height{value: value}, nil
}

func (ins *Height) Value() int {
	return ins.value
}

func (ins *Height) Equals(value *Height) bool {
	if ins == value {
		return true
	}
	return ins.value == value.Value()
}
