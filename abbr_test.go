package pinyin

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestAbbr(t *testing.T) {
	testify.Equal(t, "zg", Abbr("中国"))
	testify.Equal(t, "ryf", Abbr("阮一峰"))
}
