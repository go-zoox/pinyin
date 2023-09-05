package pinyin

import (
	"testing"

	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/testify"
)

func TestSlice(t *testing.T) {
	testify.Equal(t, strings.Join([]string{"zhong", "guo"}, ","), strings.Join(Slice("中国"), ","))
	testify.Equal(t, strings.Join([]string{"ruan", "yi", "feng"}, ","), strings.Join(Slice("阮一峰"), ","))
}
