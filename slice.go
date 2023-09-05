package pinyin

import (
	"github.com/mozillazg/go-pinyin"
)

// Slice returns the pinyin slice of the given string.
func Slice(hans string) []string {
	return pinyin.LazyConvert(hans, nil)
}
