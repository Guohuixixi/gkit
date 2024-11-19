package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	src := []int{1, 2, 3}
	dst := Map(src, func(idx int, src int) int {
		return src + 1
	})
	assert.Equal(t, []int{2, 3, 4}, dst)
}
