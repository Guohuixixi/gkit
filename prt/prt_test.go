package prt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToPtr(t *testing.T) {
	str := "测试字符串"
	ptr := ToPtr[string](str)
	assert.Equal(t, &str, ptr)
}

func TestToPrtSlice(t *testing.T) {
	src := []string{"测试字符串1", "测试字符串2"}
	assert.Equal(t, []*string{ToPtr("测试字符串1"), ToPtr("测试字符串2")}, ToPrtSlice(src))
}
