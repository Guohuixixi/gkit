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
