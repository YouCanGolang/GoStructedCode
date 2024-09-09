package main

import "fmt"

// Node 定义单向循环链表节点结构
type Node struct {
	data int
	// 指向下一个元素的指针节点
	Next *Node
}

func appendNode(head *Node, data int) *Node {
	newNode := &Node{
		data: data,
	}
	if head == nil {
		newNode.Next = newNode
		return newNode
	}

	current := head
	for current.Next != head {
		current = current.Next
	}
	current.Next = newNode
	newNode.Next = head

	return head
}

func delNode(head *Node, data int) *Node {
	if head == nil {
		return nil
	}

	current := head
	for current.Next != head {
		if current.Next.data == data {
			current.Next = current.Next.Next
			break
		}
		current = current.Next
	}
	return head
}

func printNode(head *Node) {
	if head == nil {
		return
	}

	current := head
	for current.Next != nil {
		fmt.Printf("%d -> ", current.data)

		current = current.Next
		if current == head {
			// 循环到了开头
			break
		}
	}
	fmt.Println("(回到头节点)")
}

func main() {
	var (
		head *Node
	)
	head = appendNode(head, 1)
	head = appendNode(head, 2)
	head = appendNode(head, 3)
	printNode(head)

	head = delNode(head, 2)
	printNode(head)
}
