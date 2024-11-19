package lru

// DoubleLinkedList 双向链表
type DoubleLinkedList[K comparable, V any] struct {
	key        K
	value      V
	prev, next *DoubleLinkedList[K, V]
}

// LRUCache LRU缓存
type LRUCache[K comparable, V any] struct {
	capacity   int
	cache      map[K]*DoubleLinkedList[K, V]
	head, tail *DoubleLinkedList[K, V]
}

// NewLRUCache 创建LRU缓存
func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	head, tail := &DoubleLinkedList[K, V]{}, &DoubleLinkedList[K, V]{}
	head.next = tail
	tail.prev = head
	return &LRUCache[K, V]{
		capacity: capacity,
		cache:    map[K]*DoubleLinkedList[K, V]{},
		head:     head,
		tail:     tail,
	}
}

// Get 获取缓存
func (l *LRUCache[K, V]) Get(key K) V {
	if node, ok := l.cache[key]; ok {
		l.remove(node)
		l.insert(node)
		return node.value
	}
	var v V
	return v
}

// Set 设置缓存
func (l *LRUCache[K, V]) Set(key K, value V) {
	if node, ok := l.cache[key]; ok {
		l.remove(node)
		node.value = value
		l.insert(node)
	} else {
		node = &DoubleLinkedList[K, V]{
			key:   key,
			value: value,
		}
		l.insert(node)
		l.cache[key] = node
		if len(l.cache) > l.capacity {
			node = l.tail.prev
			l.remove(node)
			delete(l.cache, node.key)
		}
	}
}

// remove 删除节点
func (l *LRUCache[K, V]) remove(node *DoubleLinkedList[K, V]) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// insert 插入节点
func (l *LRUCache[K, V]) insert(node *DoubleLinkedList[K, V]) {
	node.next = l.head.next
	node.prev = l.head
	l.head.next.prev = node
	l.head.next = node
}
