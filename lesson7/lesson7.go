package main

import "fmt"

// Tree 定义树结构
type Tree struct {
	data  int
	left  *Tree
	right *Tree
}

func (t *Tree) insert(data int) {
	newTree := &Tree{
		data: data,
	}

	if data < t.data {
		if t.left == nil {
			t.left = newTree
		} else {
			t.left.insert(data)
		}
	} else {
		if t.right == nil {
			t.right = newTree
		} else {
			t.right.insert(data)
		}
	}
}

func (t *Tree) search(data int) *Tree {
	if t == nil {
		return nil
	}

	if t.data == data {
		return t
	}

	// 递归查找左子树
	if data < t.data {
		return t.left.search(data)
	}
	// 递归查找右子树
	return t.right.search(data)
}

func (t *Tree) delete(data int) *Tree {
	if t == nil {
		return nil
	}

	if data < t.data {
		// 递归左查找
		t.left = t.left.delete(data)
	} else if data > t.data {
		// 递归右查找
		t.right = t.right.delete(data)
	} else {
		// 没有任何子节点
		if t.left == nil && t.right == nil {
			return nil
		}

		// 只有一个子节点
		if t.left == nil {
			return t.right
		}

		if t.right == nil {
			return t.left
		}

		// 有两个子节点，那么需要找到中序后继节点
		minNode := t.right.findMin()
		// 用中序后继的值替换当前节点的值
		t.data = minNode.data
		// 删除中序后继节点
		t.right = t.right.delete(minNode.data)
	}
	return t
}

// 查找最小节点（中序后继）
func (t *Tree) findMin() *Tree {
	if t.left == nil {
		return t
	}
	return t.left.findMin()
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

func (t *Tree) printTree() {
	if t == nil {
		return
	}
	if t.left != nil {
		fmt.Printf("left: %d, ", t.left.data)
	}
	if t.right != nil {
		fmt.Printf("right-> %d\n", t.right.data)
	}
	if t.left == nil && t.right == nil {
		return
	}
	t.left.printTree()
	t.right.printTree()
}

func main() {
	tree := &Tree{data: 8}
	tree.insert(3)
	tree.insert(10)
	tree.insert(1)
	tree.insert(6)
	tree.insert(14)
	tree.insert(4)
	tree.insert(7)
	tree.insert(13)
	tree.printTree()
	fmt.Println("\n中序遍历结果:")
	tree.inOrder()

	fmt.Println("\n\n查找节点 6")
	node := tree.search(6)
	if node != nil {
		fmt.Printf("找到节点: %d\n", node.data)
	} else {
		fmt.Println("节点不存在")
	}

	fmt.Println("\n删除节点 13 后的中序遍历结果:")
	tree = tree.delete(13)
	tree.inOrder()

	fmt.Println("\n删除节点 10 后的中序遍历结果:")
	tree = tree.delete(10)
	tree.inOrder()

	fmt.Println("\n删除节点 6 后的中序遍历结果:")
	tree = tree.delete(6)
	tree.inOrder()
}
