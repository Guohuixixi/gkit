package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapSet(t *testing.T) {
	set := NewMapSet[int](10)

	// 测试用例1：向集合中添加元素
	set.Add(1)
	set.Add(2)
	set.Add(3)
	if set.Size() != 3 {
		t.Errorf("预期大小为3，但得到%d", set.Size())
	}

	// 测试用例2：从集合中获取所有元素
	expected := []int{1, 2, 3}
	got := set.GetAll()
	if len(got) != len(expected) {
		t.Errorf("预期为%v，但得到%v", expected, got)
	}

	// 测试用例3：从集合中移除一个元素
	set.Remove(2)
	if set.Size() != 2 {
		t.Errorf("预期大小为2，但得到%d", set.Size())
	}

	// 测试用例4：检查集合中是否存在某个元素
	if !set.Contains(1) {
		t.Errorf("预期1存在于集合中")
	}

	// 测试用例5：检查集合中是否不存在某个元素
	if set.Contains(2) {
		t.Errorf("预期2不存在于集合中")
	}

	// 测试用例6：获取集合的大小
	if set.Size() != 2 {
		t.Errorf("预期大小为2，但得到%d", set.Size())
	}

	// 测试用例7：清空集合
	set.Clear()
	if set.Size() != 0 {
		t.Errorf("预期大小为0，但得到%d", set.Size())
	}

	// 测试用例8：集合的字符串表示
	expectedStr := "{}"
	gotStr := set.String()
	if gotStr != expectedStr {
		t.Errorf("预期为%s，但得到%s", expectedStr, gotStr)
	}
}

func TestMapSet_Add(t *testing.T) {
	values := []int{1, 2, 3, 1}
	set := NewMapSet[int](10)
	t.Run("Add", func(t *testing.T) {
		for _, val := range values {
			set.Add(val)
		}
		assert.Equal(t, set.m, map[int]struct{}{
			1: struct{}{},
			2: struct{}{},
			3: struct{}{},
		})
	})
}
