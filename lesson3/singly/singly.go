package main

import (
	"fmt"
)

// Node 定义单向列表的节点结构
type Node struct {
	// 数据
	data int
	// 指向下一个节点
	next *Node
}

func appendNode(head *Node, data int) *Node {
	newNode := &Node{
		data: data,
	}

	// 第一个元素
	if head == nil {
		return newNode
	}

	current := head

	// 如果当前元素是有指向的，那么使用下一个
	for current.next != nil {
		current = head.next
	}

	// 指向新的元素
	current.next = newNode

	return head
}

func deleteNode(head *Node, data int) *Node {
	if head == nil {
		return nil
	}

	current := head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			break
		}
		current = current.next
	}

	return head
}

func printNode(head *Node) {
	current := head

	// 读取所有元素展示
	for current != nil {
		fmt.Print(current.data, " -> ")
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

	head = deleteNode(head, 2)
	printNode(head)
}
