package main

import "fmt"

// Node 定义双向列表的节点结构
type Node struct {
	// 数据
	data int
	// 指向上一个节点
	prev *Node
	// 指向下一个节点
	next *Node
}

func appendNode(head *Node, data int) *Node {
	newNode := &Node{
		data: data,
	}
	if head == nil {
		return newNode
	}

	current := head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	newNode.prev = current
	return head
}

func delNode(head *Node, data int) *Node {
	if head == nil {
		return nil
	}

	current := head
	for current.next != nil {
		// 匹配到
		if current.next.data == data {
			// 更新当前的后续指向
			current.next = current.next.next
			if current.next.next != nil {
				// 更新新的前序指向，将它们链接起来
				current.next.next.prev = current
			}
			break
		}
		current = current.next
	}
	return head
}

func printNode(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d <-> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	var (
		head *Node
	)
	head = appendNode(head, 1)
	head = appendNode(head, 2)
	head = appendNode(head, 3)
	printNode(head)

	delNode(head, 2)
	printNode(head)
}
