package binary_tree

import (
	"fmt"
	"math"
)

type node[T any] struct {
	Left   *node[T]
	Right  *node[T]
	Parent *node[T]
	Value  T
}

func NewNode[T any](elem T) *node[T] {
	return &node[T]{
		Value:  elem,
		Left:   nil,
		Right:  nil,
		Parent: nil,
	}
}

func getSpaceCount(height int) int {
	return int(3*math.Pow(2, float64(height-2)) - 1)
}

func getSlashCount(height int) int {
	if height <= 3 {
		return height - 1
	}
	return int(3*math.Pow(2, float64(height-3)) - 1)
}

func getHeight[T any](root *node[T]) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(getHeight(root.Left)), float64(getHeight(root.Right))))
}

func nodePrint[T any](prefix string, n *node[T], isLeft bool) string {
	if n == nil {
		return ""
	}
	//fmt.Println(prefix)
	prefix = prefix + func() string {
		if isLeft {
			return fmt.Sprintf("|--%v\n", n.Value)
		}
		return fmt.Sprintf("\\--%v\n", n.Value)
	}()
	nodePrint(func() string {
		if isLeft {
			return prefix + "|   "
		}
		return prefix + "    "
	}(), n.Left, true)
	nodePrint(func() string {
		if isLeft {
			return prefix + "|   "
		}
		return prefix + "    "
	}(), n.Right, false)
	return prefix
}

type BinaryTree[T any] struct {
	Root *node[T]
}

func (b *BinaryTree[T]) String() string {
	return nodePrint("", b.Root, false)
}

func (b *BinaryTree[T]) IsEmpty() bool {
	return b.Root == nil
}

func inOrder[T any](node *node[T], iter chan T) {
	if node == nil {
		return
	}
	inOrder(node.Left, iter)
	iter <- node.Value
	inOrder(node.Right, iter)
}

func (b *BinaryTree[T]) Iter() <-chan T {
	iter := make(chan T)
	go func() {
		inOrder(b.Root, iter)
		close(iter)
	}()
	return iter
}

func (b *BinaryTree[T]) Insert(elem T, less func(a, b T) bool) {
	if b.Root == nil {
		b.Root = NewNode(elem)
	} else {
		current := b.Root
		parent := current
		for current != nil {
			parent = current
			if less(elem, current.Value) {
				current = current.Left
			} else {
				current = current.Right
			}
		}
		current = NewNode(elem)
		current.Parent = parent
		if less(current.Value, parent.Value) {
			parent.Left = current
		} else {
			parent.Right = current
		}
	}
}
