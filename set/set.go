package set

import (
	"fmt"
	"strings"
)

// Set 集合
type Set[T comparable] interface {
	// Add 添加元素
	Add(val T)
	// GetAll 获取集合所有元素
	GetAll() []T
	// Remove 删除元素
	Remove(key T)
	// Contains 判断元素是否存在
	Contains(key T) bool
	// Size 获取集合大小
	Size() int
	// Clear 清空集合
	Clear()
	// String 返回集合字符串表示
	String() string
}

// MapSet 基于Map实现的Set集合
type MapSet[T comparable] struct {
	m map[T]struct{}
}

// NewMapSet 新建MapSet
func NewMapSet[T comparable](size int) *MapSet[T] {
	return &MapSet[T]{
		m: make(map[T]struct{}, size),
	}
}

// Add 添加元素
func (s *MapSet[T]) Add(val T) {
	s.m[val] = struct{}{}
}

// GetAll 方法返回的元素顺序不固定
func (s *MapSet[T]) GetAll() []T {
	res := make([]T, 0, len(s.m))
	for key := range s.m {
		res = append(res, key)
	}
	return res
}

// Remove 删除元素
func (s *MapSet[T]) Remove(key T) {
	delete(s.m, key)
}

// Contains 判断元素是否存在
func (s *MapSet[T]) Contains(key T) bool {
	_, ok := s.m[key]
	return ok
}

// Size 获取集合大小
func (s *MapSet[T]) Size() int {
	return len(s.m)
}

func (s *MapSet[T]) Clear() {
	s.m = make(map[T]struct{})
}

// String 返回集合字符串表示
func (s *MapSet[T]) String() string {
	var vals []string
	for val := range s.m {
		vals = append(vals, fmt.Sprintf("%v", val))
	}
	return fmt.Sprintf("{%v}", strings.Join(vals, ", "))
}
