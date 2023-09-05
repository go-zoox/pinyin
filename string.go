package pinyin

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/mozillazg/go-pinyin"
)

// String returns the pinyin string of the given string.
func String(hans string) string {
	parts := pinyin.LazyConvert(hans, nil)
	return strings.Join(parts, "")
}
