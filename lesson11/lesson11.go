package main

import (
	"fmt"
)

// item 定义节点结构
type item struct {
	// 任务
	task string
	// 优先级
	priority int
}

// PriorityQueue 定义优先队列结构
type PriorityQueue struct {
	items []*item
}

// Insert 插入元素
func (pq *PriorityQueue) Insert(task string, priority int) {
	pq.items = append(pq.items, &item{
		task:     task,
		priority: priority,
	})
	pq.siftUp(len(pq.items) - 1)
}

// 上滤操作，维持最大堆性质
func (pq *PriorityQueue) siftUp(index int) {
	parent := (index - 1) / 2
	if parent >= 0 && pq.items[parent].priority < pq.items[index].priority {
		pq.items[parent], pq.items[index] = pq.items[index], pq.items[parent]
		pq.siftUp(parent)
	}
}

// Remove 删除堆顶元素（优先级最高）
func (pq *PriorityQueue) Remove() *item {
	if len(pq.items) == 0 {
		return nil
	}

	// 将堆顶元素和最后一个元素交换，并删除堆顶元素
	top := pq.items[0]
	pq.items[0] = pq.items[len(pq.items)-1]
	pq.items = pq.items[:len(pq.items)-1]

	// 下滤操作，恢复堆序
	pq.siftDown(0)

	return top
}

// 下滤操作，维持最大堆性质
func (pq *PriorityQueue) siftDown(index int) {
	left := 2*index + 1
	right := 2*index + 2
	largest := index

	if left < len(pq.items) && pq.items[left].priority > pq.items[largest].priority {
		largest = left
	}
	if right < len(pq.items) && pq.items[right].priority > pq.items[largest].priority {
		largest = right
	}

	if largest != index {
		pq.items[index], pq.items[largest] = pq.items[largest], pq.items[index]
		pq.siftDown(largest)
	}
}

func (pq *PriorityQueue) print() {
	for _, v := range pq.items {
		fmt.Printf("任务 %s, 优先级 %d\n", v.task, v.priority)
	}
}

func main() {
	pq := &PriorityQueue{}
	pq.Insert("A", 5)
	pq.Insert("B", 3)
	pq.Insert("C", 8)
	pq.Insert("D", 2)
	pq.Insert("E", 7)

	fmt.Println("优先队列: ")
	pq.print()

	fmt.Println("删除优先级最高的元素: ", pq.Remove())
	fmt.Println("优先队列: ")
	pq.print()

	fmt.Println("删除优先级最高的元素: ", pq.Remove())
	fmt.Println("优先队列: ")
	pq.print()
}
