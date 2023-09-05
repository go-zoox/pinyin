package pinyin

import (
	"testing"

	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/testify"
)

func TestString(t *testing.T) {
	testify.Equal(t, strings.Join([]string{"zhong", "guo"}, ""), String("中国"))
	testify.Equal(t, strings.Join([]string{"ruan", "yi", "feng"}, ""), String("阮一峰"))
}
