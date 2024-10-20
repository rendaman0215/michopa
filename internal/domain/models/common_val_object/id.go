package commonvalobject

import (
	"fmt"
	"regexp"
	"unicode/utf8"

	"github.com/rendaman0215/michopa/internal/err"
)

type ID struct {
	value string
}

// valueフィールドのゲッター
func (ins *ID) Value() string {
	return ins.value
}

// 同一性検証
func (ins *ID) Equals(value *ID) bool {
	if ins == value {
		return true
	}
	return ins.value == value.Value()
}

func NewId(value string) (*ID, error) {
	// フィールドの長さ
	const LENGTH int = 36
	// UUIDの正規表現
	const REGEXP string = "([0-9a-f]{8})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{12})"
	// 引数の文字数チェック
	if utf8.RuneCountInString(value) != LENGTH {
		return nil, err.NewInternalError(fmt.Sprintf("カテゴリIDの長さは%d文字でなければなりません。", LENGTH))
	}
	// 引数の正規表現(UUID)チェック
	if !regexp.MustCompile(REGEXP).Match([]byte(value)) {
		return nil, err.NewInternalError("カテゴリIDはUUIDの形式でなければなりません。")
	}
	return &ID{value: value}, nil
}
