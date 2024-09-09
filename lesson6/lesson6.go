package main

import "fmt"

// Tree 定义二叉树结构
type Tree struct {
	// 数据
	data int
	// 左节点
	left *Tree
	// 右节点
	right *Tree
}

func (n *Tree) insert(data int) {
	newNode := &Tree{
		data: data,
	}

	// 奇数处理
	if data%2 != 0 {
		if n.left == nil {
			n.left = newNode
			return
		}

		if n.right == nil {
			n.right = newNode
			return
		}

		n.left.insert(data)
		return
	}

	// 偶数处理
	if n.right == nil {
		n.right = newNode
		return
	}

	if n.left == nil {
		n.left = newNode
		return
	}
	n.right.insert(data)
}

func (n *Tree) printTree() {
	if n == nil {
		return
	}
	if n.left != nil {
		fmt.Printf("%d, ", n.left.data)
	}
	if n.right != nil {
		fmt.Printf("%d\n", n.right.data)
	}
	if n.left == nil && n.right == nil {
		return
	}
	n.left.printTree()
	n.right.printTree()
}

func main() {
	node := &Tree{
		data: 1,
	}
	node.insert(2)
	node.insert(3)
	node.insert(4)
	node.insert(5)
	node.insert(6)
	node.insert(7)
	node.printTree()
}
