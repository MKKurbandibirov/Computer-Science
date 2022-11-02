package binary_tree

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

type BinaryTree[T any] struct {
	Root *node[T]
}

func (b *BinaryTree[T]) printNodeInternal() {

}

//func (b *BinaryTree[T]) String() string {
//
//}

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
