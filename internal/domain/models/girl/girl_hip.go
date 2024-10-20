package girl

import "github.com/rendaman0215/michopa/internal/err"

type Hip struct {
	value int
}

func NewHip(value int) (*Hip, error) {
	// ヒップが70cm以上、120cm以下であるかどうか
	if value < 70 || value > 120 {
		return nil, err.NewInternalError("ヒップは70cm以上、120cm以下でなければなりません。")
	}

	return &Hip{value: value}, nil
}

func (ins *Hip) Value() int {
	return ins.value
}

func (ins *Hip) Equals(value *Hip) bool {
	if ins == value {
		return true
	}
	return ins.value == value.Value()
}
