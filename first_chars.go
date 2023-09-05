package pinyin

import (
	"github.com/go-zoox/core-utils/array"
	"github.com/mozillazg/go-pinyin"
)

// FirstChars returns the first char of every pinyin of the given string.
func FirstChars(hans string) string {
	parts := pinyin.LazyConvert(hans, nil)
	firstChars := array.Map(parts, func(v string, i int) byte {
		return v[0]
	})
	return string(firstChars)
}
