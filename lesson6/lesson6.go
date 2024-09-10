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

func (t *Tree) insert(data int) {
	newTree := &Tree{
		data: data,
	}

	// 奇数处理
	if data%2 != 0 {
		if t.left == nil {
			t.left = newTree
			return
		}

		if t.right == nil {
			t.right = newTree
			return
		}

		t.left.insert(data)
		return
	}

	// 偶数处理
	if t.right == nil {
		t.right = newTree
		return
	}

	if t.left == nil {
		t.left = newTree
		return
	}
	t.right.insert(data)
}

func (t *Tree) printTree() {
	if t == nil {
		return
	}
	if t.left != nil {
		fmt.Printf("%d, ", t.left.data)
	}
	if t.right != nil {
		fmt.Printf("%d\n", t.right.data)
	}
	if t.left == nil && t.right == nil {
		return
	}
	t.left.printTree()
	t.right.printTree()
}

func main() {
	tree := &Tree{
		data: 1,
	}
	tree.insert(2)
	tree.insert(3)
	tree.insert(4)
	tree.insert(5)
	tree.insert(6)
	tree.insert(7)
	tree.printTree()
}
