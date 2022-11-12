package heap

import (
	"fmt"
	"strings"
)

const (
	MAX = iota
	MIN
)

type Type int

type Heap[T any] struct {
	arr      []T
	less     func(x, y T) bool
	more     func(x, y T) bool
	heapType Type
	size     int
}

func NewHeap[T any](less, more func(a, b T) bool, heapType Type) *Heap[T] {
	return &Heap[T]{
		arr:      nil,
		less:     less,
		more:     more,
		heapType: heapType,
		size:     0,
	}
}

func (h *Heap[T]) Iter() <-chan T {
	iter := make(chan T)
	go func() {
		for i := h.Size() - 1; i >= 0; i-- {
			iter <- h.GetRoot()
		}
		close(iter)
		h.size = len(h.arr)
		h.Build(h.arr)
	}()
	return iter
}

func (h *Heap[T]) stringHelper(prefix string, i int, isLeft bool, sb *strings.Builder) {
	if i < h.Size() {
		sb.WriteString(prefix + func() string {
			if isLeft {
				return fmt.Sprintf("|--%v\n", h.arr[i])
			}
			return fmt.Sprintf("\\--%v\n", h.arr[i])
		}())
		h.stringHelper(prefix+func() string {
			if isLeft {
				return "|   "
			}
			return "    "
		}(), i*2+2, true, sb)
		h.stringHelper(prefix+func() string {
			if isLeft {
				return "|   "
			}
			return "    "
		}(), i*2+1, false, sb)
	}
}

func (h *Heap[T]) String() string {
	var sb strings.Builder
	h.stringHelper("", 0, false, &sb)
	return sb.String()
}

func (h *Heap[T]) Size() int {
	return h.size
}

func (h *Heap[T]) IsEmpty() bool {
	return h.Size() == 0
}

func (h *Heap[T]) Add(elem T) {
	h.arr = append(h.arr, elem)
	i := h.Size() - 1
	parent := (i - 1) / 2
	comp := h.less
	if h.heapType == MIN {
		comp = h.more
	}
	for i > 0 && comp(h.arr[parent], h.arr[i]) {
		h.arr[parent], h.arr[i] = h.arr[i], h.arr[parent]
		i = parent
		parent = (i - 1) / 2
	}
	h.size++
}

func (h *Heap[T]) Build(elems []T) {
	h.arr = elems
	for i := h.Size() / 2; i >= 0; i-- {
		h.Heapify(i)
	}
}

func (h *Heap[T]) Heapify(i int) {
	var left, right, large int
	comp := h.more
	if h.heapType == MIN {
		comp = h.less
	}
	for {
		left = 2*i + 1
		right = 2*i + 2
		large = i
		if left < h.Size() && comp(h.arr[left], h.arr[large]) {
			large = left
		}
		if right < h.Size() && comp(h.arr[right], h.arr[large]) {
			large = right
		}
		if large == i {
			break
		}
		h.arr[i], h.arr[large] = h.arr[large], h.arr[i]
		i = large
	}
}

func (h *Heap[T]) GetRoot() T {
	result := h.arr[0]
	h.arr[0], h.arr[h.Size()-1] = h.arr[h.Size()-1], h.arr[0]
	h.size--
	h.Heapify(0)
	return result
}
