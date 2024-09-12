package main

import "fmt"

// MaxHeap 定义堆结构
type MaxHeap struct {
	heap []int
}

// Insert 插入新元素
func (h *MaxHeap) Insert(val int) {
	h.heap = append(h.heap, val)
	h.siftUp(len(h.heap) - 1)
}

// 上滤操作，维持最大堆性质
func (h *MaxHeap) siftUp(index int) {
	parent := (index - 1) / 2
	if parent >= 0 && h.heap[parent] < h.heap[index] {
		h.heap[parent], h.heap[index] = h.heap[index], h.heap[parent]
		h.siftUp(parent)
	}
}

// 删除堆顶元素
func (h *MaxHeap) Remove() int {
	if len(h.heap) == 0 {
		return -1
	}

	top := h.heap[0]
	// 最后一个元素移动到这里
	h.heap[0] = h.heap[len(h.heap)-1]
	// 去除最后一个元素
	h.heap = h.heap[:len(h.heap)-1]
	h.siftDown(0)
	return top
}

// 下滤操作，维持最大堆性质
func (h *MaxHeap) siftDown(index int) {
	left := 2*index + 1
	right := 2*index + 2
	largest := index

	if left < len(h.heap) && h.heap[left] > h.heap[largest] {
		largest = left
	}
	if right < len(h.heap) && h.heap[right] > h.heap[largest] {
		largest = right
	}

	if largest != index {
		h.heap[index], h.heap[largest] = h.heap[largest], h.heap[index]
		h.siftDown(largest)
	}
}

func main() {
	h := &MaxHeap{}
	h.Insert(9)
	h.Insert(5)
	h.Insert(6)
	h.Insert(2)
	h.Insert(3)
	h.Insert(4)

	fmt.Println("堆:", h.heap)

	fmt.Println("删除堆顶:", h.Remove())
	fmt.Println("堆:", h.heap)
}
