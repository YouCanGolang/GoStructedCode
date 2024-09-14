package main

import (
	"fmt"
)

// Entry 定义键值对结构
type Entry struct {
	key   string
	value string
	next  *Entry
}

// HashTable 定义哈希表结构
type HashTable struct {
	table []*Entry
	size  int
}

// NewHashTable 创建一个哈希表
func NewHashTable(size int) *HashTable {
	return &HashTable{
		table: make([]*Entry, size),
		size:  size,
	}
}

// Hash 函数计算哈希值
func (h *HashTable) Hash(key string) int {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}
	return hash % h.size
}

// Put 插入键值对到哈希表
func (h *HashTable) Put(key string, value string) {
	index := h.Hash(key)
	entry := h.table[index]

	// 如果该位置为空，直接插入
	if entry == nil {
		h.table[index] = &Entry{key: key, value: value}
		return
	}

	// 解决碰撞，检查链表中是否已经存在该键
	for entry != nil {
		if entry.key == key {
			entry.value = value // 更新值
			return
		}
		if entry.next == nil {
			break
		}
		entry = entry.next
	}

	// 将新的键值对插入链表末尾
	entry.next = &Entry{key: key, value: value}
}

// Get 根据键查找值
func (h *HashTable) Get(key string) (string, bool) {
	index := h.Hash(key)
	entry := h.table[index]

	for entry != nil {
		if entry.key == key {
			return entry.value, true
		}
		entry = entry.next
	}
	return "", false
}

// Remove 删除键值对
func (h *HashTable) Remove(key string) {
	index := h.Hash(key)
	entry := h.table[index]

	if entry == nil {
		return
	}

	// 如果第一个元素就是要删除的键
	if entry.key == key {
		h.table[index] = entry.next
		return
	}

	// 遍历链表寻找要删除的键
	prev := entry
	entry = entry.next
	for entry != nil {
		if entry.key == key {
			prev.next = entry.next
			return
		}
		prev = entry
		entry = entry.next
	}
}

func main() {
	// 创建哈希表
	hashTable := NewHashTable(10)

	// 插入键值对
	hashTable.Put("A", "Apple")
	hashTable.Put("B", "Banana")
	hashTable.Put("C", "Cherry")

	// 查找键值对
	value, found := hashTable.Get("B")
	if found {
		fmt.Println("Key B:", value)
	} else {
		fmt.Println("Key B not found")
	}

	// 删除键值对
	hashTable.Remove("B")
	fmt.Println("remove B")

	value, found = hashTable.Get("B")
	if found {
		fmt.Println("Key B:", value)
	} else {
		fmt.Println("Key B not found")
	}
}
