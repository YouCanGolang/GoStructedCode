package main

import "fmt"

type Queue struct {
	elements []int
}

func (q *Queue) Enqueue(ele int) {
	q.elements = append(q.elements, ele)
}

func (q *Queue) Dequeue() int {
	if len(q.elements) == 0 {
		fmt.Println("队列为空")
		return -1
	}
	front := q.elements[0]
	// 移除队首元素
	q.elements = q.elements[1:]
	return front
}

func (q *Queue) Peek() int {
	if len(q.elements) == 0 {
		fmt.Println("队列为空")
		return -1
	}
	return q.elements[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func main() {
	queue := &Queue{}
	fmt.Println("queue is empty-> ", queue.IsEmpty())
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	queue.Dequeue()
	fmt.Println("queue peek-> ", queue.Peek())
}
