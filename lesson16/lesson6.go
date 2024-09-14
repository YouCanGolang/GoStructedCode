package main

import "fmt"

// HashTable 定义哈希表结构
type HashTable struct {
	table []string
	size  int
}

// NewHashTable 创建一个哈希表
func NewHashTable(size int) *HashTable {
	return &HashTable{
		table: make([]string, size),
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

	// 如果发生冲突，进行线性探测
	for h.table[index] != "" {
		index = (index + 1) % h.size // 移动到下一个位置
	}

	h.table[index] = value
}

// Get 根据键查找值
func (h *HashTable) Get(key string) (string, bool) {
	index := h.Hash(key)

	// 线性探测查找值
	for h.table[index] != "" {
		return h.table[index], true
	}
	return "", false
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
		fmt.Println("Key B:", value) // 输出 Key B: Banana
	} else {
		fmt.Println("Key B not found")
	}
}
