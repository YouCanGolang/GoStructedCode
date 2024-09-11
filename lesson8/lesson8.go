package main

import (
	"fmt"
)

// Tree 定义树结构
type Tree struct {
	data   int
	left   *Tree
	right  *Tree
	height int
}

// 获取节点的高度
func (t *Tree) getHeight() int {
	if t == nil {
		return 0
	}
	return t.height
}

// 计算平衡因子
func (t *Tree) getBalanceFactor() int {
	if t == nil {
		return 0
	}
	return t.left.getHeight() - t.right.getHeight()
}

// 更新节点的高度
func (t *Tree) updateHeight() {
	leftHeight := t.left.getHeight()
	rightHeight := t.right.getHeight()
	if leftHeight > rightHeight {
		t.height = leftHeight + 1
	} else {
		t.height = rightHeight + 1
	}
}

// 左旋操作
func (t *Tree) leftRotate() *Tree {
	newRoot := t.right
	t.right = newRoot.left
	newRoot.left = t

	// 更新高度
	t.updateHeight()
	newRoot.updateHeight()

	return newRoot
}

// 右旋操作
func (t *Tree) rightRotate() *Tree {
	newRoot := t.left
	t.left = newRoot.right
	newRoot.right = t

	// 更新高度
	t.updateHeight()
	newRoot.updateHeight()

	return newRoot
}

// 插入节点
func (t *Tree) insert(data int) *Tree {
	if t == nil {
		return &Tree{data: data, height: 1}
	}

	if data < t.data {
		t.left = t.left.insert(data)
	} else if data > t.data {
		t.right = t.right.insert(data)
	} else {
		// 不允许重复节点
		return t
	}

	// 更新当前节点的高度
	t.updateHeight()

	// 检查是否需要旋转
	balance := t.getBalanceFactor()

	// 左子树过高，进行右旋
	if balance > 1 {
		if data < t.left.data {
			return t.rightRotate()
		}
		// 左右情况，先左旋后右旋
		if data > t.left.data {
			t.left = t.left.leftRotate()
			return t.rightRotate()
		}
	}

	// 右子树过高，进行左旋
	if balance < -1 {
		if data > t.right.data {
			return t.leftRotate()
		}
		// 右左情况，先右旋后左旋
		if data < t.right.data {
			t.right = t.right.rightRotate()
			return t.leftRotate()
		}
	}

	return t
}

// 中序遍历
func (t *Tree) inOrder() {
	if t == nil {
		return
	}
	t.left.inOrder()
	fmt.Printf("%d ", t.data)
	t.right.inOrder()
}

func main() {
	root := &Tree{data: 10}
	root = root.insert(20)
	root = root.insert(30)
	root = root.insert(40)
	root = root.insert(50)
	root = root.insert(25)

	fmt.Println("中序遍历结果：")
	root.inOrder()
}
