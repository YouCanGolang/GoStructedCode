package main

import (
	"fmt"
)

const (
	RED   = true
	BLACK = false
)

// RBTree 定义红黑树节点结构
type RBTree struct {
	data  int
	left  *RBTree
	right *RBTree
	color bool // true表示红色，false表示黑色
}

// 旋转操作
// 左旋
func leftRotate(t *RBTree) *RBTree {
	newRoot := t.right
	t.right = newRoot.left
	newRoot.left = t
	newRoot.color = t.color
	t.color = RED
	return newRoot
}

// 右旋
func rightRotate(t *RBTree) *RBTree {
	newRoot := t.left
	t.left = newRoot.right
	newRoot.right = t
	newRoot.color = t.color
	t.color = RED
	return newRoot
}

// 颜色翻转
func flipColors(t *RBTree) {
	t.color = RED
	t.left.color = BLACK
	t.right.color = BLACK
}

// 插入节点
func (t *RBTree) insert(data int) *RBTree {
	if t == nil {
		return &RBTree{data: data, color: RED}
	}

	if data < t.data {
		t.left = t.left.insert(data)
	} else if data > t.data {
		t.right = t.right.insert(data)
	}

	// 修复红黑树性质
	if isRed(t.right) && !isRed(t.left) {
		t = leftRotate(t)
	}
	if isRed(t.left) && isRed(t.left.left) {
		t = rightRotate(t)
	}
	if isRed(t.left) && isRed(t.right) {
		flipColors(t)
	}

	return t
}

// 判断节点是否是红色
func isRed(t *RBTree) bool {
	if t == nil {
		return false
	}
	return t.color == RED
}

// 中序遍历
func (t *RBTree) inOrder() {
	if t == nil {
		return
	}
	t.left.inOrder()
	fmt.Printf("%d ", t.data)
	t.right.inOrder()
}

func main() {
	root := &RBTree{data: 10, color: BLACK}
	root = root.insert(20)
	root = root.insert(30)
	root = root.insert(40)
	root = root.insert(50)
	root = root.insert(25)

	fmt.Println("中序遍历结果：")
	root.inOrder()
}
