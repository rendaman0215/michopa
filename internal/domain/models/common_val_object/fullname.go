package commonvalobject

import (
	"strings"

	"github.com/rendaman0215/michopa/internal/err"
)

type Fullname struct {
	firstname string
	lastname  string
}

// getter
func (ins *Fullname) Firstname() string {
	return ins.firstname
}
func (ins *Fullname) Lastname() string {
	return ins.lastname
}

// コンストラクタ
func NewFullname(firstname, lastname string) (*Fullname, error) {
	if len(firstname) == 0 || len(lastname) == 0 {
		return nil, err.NewInternalError("名前は空であってはなりません。")
	}
	return &Fullname{
		firstname: strings.TrimSpace(firstname),
		lastname:  strings.TrimSpace(lastname),
	}, nil
}

// 同一性検証
func (ins *Fullname) Equals(value *Fullname) bool {
	if ins == value {
		return true
	}
	return ins.firstname == value.Firstname() && ins.lastname == value.Lastname()
}
