package structure

import (
	"errors"
	"fmt"
)

type treeNode struct {
	val         int
	left, right *treeNode
}

type BinarySearchTree struct {
	root *treeNode
	size int
}

func (b *BinarySearchTree) Size() int {
	return b.size
}

func (b *BinarySearchTree) IsEmpty() bool {
	return b.size == 0
}

// Add 向bst中添加元素
func (b *BinarySearchTree) Add(val int) {
	if b.root == nil {
		b.root = &treeNode{val, nil, nil}
	} else {
		b.root = b.add(b.root, val)
	}
}

func (b *BinarySearchTree) add(node *treeNode, val int) *treeNode {

	if node == nil {
		b.size++
		return &treeNode{val, nil, nil}
	}

	if node.val > val {
		node.left = b.add(node.left, val)
	} else if node.val < val {
		node.right = b.add(node.right, val)
	}

	return node
}

// Contains 二分搜索树中是否包含指定元素
func (b *BinarySearchTree) Contains(val int) bool {
	return b.contains(b.root, val)
}

func (b *BinarySearchTree) contains(node *treeNode, val int) bool {
	if node.val == val {
		return true
	} else if node.val > val {
		return b.contains(node.left, val)
	} else {
		return b.contains(node.right, val)
	}

}

// PreOrder 前序遍历
func (b *BinarySearchTree) PreOrder() {
	b.preOrder(b.root)
}

// PreOrderNR 非递归前序遍历
func (b *BinarySearchTree) PreOrderNR() {
	if b.root == nil {
		return
	}
	stack := make([]*treeNode, 0)
	stack = append(stack, b.root)
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 访问该节点
		fmt.Println(cur.val)

		if cur.right != nil {
			stack = append(stack, cur.right)
		}

		if cur.left != nil {
			stack = append(stack, cur.left)
		}
	}
}

func (b *BinarySearchTree) preOrder(node *treeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.val)
	b.preOrder(node.left)
	b.preOrder(node.right)
}

// InOrder 中序遍历
func (b *BinarySearchTree) InOrder() {
	b.inOrder(b.root)
}

func (b *BinarySearchTree) inOrder(node *treeNode) {
	if node == nil {
		return
	}

	b.inOrder(node.left)
	fmt.Println(node.val)
	b.inOrder(node.right)
}

// LevelOrder 层序遍历
func (b *BinarySearchTree) LevelOrder() {
	if b.root == nil {
		return
	}

	q := make([]*treeNode, 0)
	q = append(q, b.root)

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		fmt.Println(cur.val)
		q = append(q, cur.left)
		q = append(q, cur.right)
	}
}

// Minimum 树中的最小数
func (b *BinarySearchTree) Minimum() (int, error) {
	if b.size == 0 {
		return 0, errors.New("empty")
	}
	node := b.minimum(b.root)
	return node.val, nil
}

func (b *BinarySearchTree) minimum(node *treeNode) *treeNode {
	if node.left == nil {
		return node
	}
	return b.minimum(node.left)
}

// Maximum 树中的最大值
func (b *BinarySearchTree) Maximum() (int, error) {
	if b.size == 0 {
		return 0, errors.New("empty")
	}
	node := b.minimum(b.root)
	return node.val, nil
}

func (b *BinarySearchTree) maximum(node *treeNode) *treeNode {
	if node.right == nil {
		return node
	}
	return b.minimum(node.right)
}

// RemoveMin 删除最小值
func (b *BinarySearchTree) RemoveMin() int {
	minimum, _ := b.Minimum()
	b.root = b.removeMin(b.root)

	return minimum
}

func (b *BinarySearchTree) removeMin(node *treeNode) *treeNode {
	if node.left == nil {
		rightNode := node.right
		node.right = nil

		b.size--
		return rightNode
	}
	node.left = b.removeMin(node.left)
	return node
}

// RemoveMax 删除最大值
func (b *BinarySearchTree) RemoveMax() int {
	maximum, _ := b.Maximum()
	b.root = b.removeMax(b.root)

	return maximum
}

func (b *BinarySearchTree) removeMax(node *treeNode) *treeNode {
	if node.right == nil {
		leftNode := node.left
		node.left = nil

		b.size--
		return leftNode
	}
	node.right = b.removeMax(node.right)
	return node
}

func (b *BinarySearchTree) Remove(val int) {
	b.root = b.remove(b.root, val)
}

func (b *BinarySearchTree) remove(node *treeNode, val int) *treeNode {
	if node == nil {
		return nil
	}

	if node.val > val {
		node.left = b.remove(node.left, val)
		return node
	} else if node.val < val {
		node.right = b.remove(node.right, val)
		return node
	} else { // node.val == val
		if node.left == nil { // 左子树为空 返回右子树
			rightNode := node.right
			node.right = nil
			b.size--
			return rightNode
		} else if node.right == nil { // 右子树为空 返回左子树
			leftNode := node.left
			node.left = nil
			b.size--
			return leftNode
		} else { // 左右子树都不为空
			// 在右子树中找到node的后继节点
			successor := b.minimum(node.right)

			// 右子树为删除右子树中最小值的根节点
			successor.right = b.removeMin(node.right)
			// 左子树就为node的左子树
			successor.left = node.left
			node.left = nil
			node.right = nil

			return successor
		}
	}
}
