package pinyin

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestFirstChars(t *testing.T) {
	testify.Equal(t, "zg", FirstChars("中国"))
	testify.Equal(t, "ryf", FirstChars("阮一峰"))
}
