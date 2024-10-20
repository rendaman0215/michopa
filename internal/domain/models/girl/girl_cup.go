package girl

import (
	"regexp"

	"github.com/rendaman0215/michopa/internal/err"
)

type Cup struct {
	value string
}

func NewCup(value string) (*Cup, error) {
	// カップがAからMまでのアルファベットであるかどうか正規表現でチェック
	validCup := regexp.MustCompile("^[A-M]$").MatchString(value)
	if !validCup {
		return nil, err.NewInternalError("カップはAからMまでのアルファベットでなければなりません。")
	}

	return &Cup{value: value}, nil
}

func (ins *Cup) Value() string {
	return ins.value
}

func (ins *Cup) Equals(value *Cup) bool {
	if ins == value {
		return true
	}
	return ins.value == value.Value()
}
