package array

import (
	"errors"
	"fmt"
	"strings"
)

type Array[T any] struct {
	arr []T
}

func (a *Array[T]) String() string {
	var builder strings.Builder
	builder.WriteString("[")
	i := 0
	for ; i < len(a.arr)-1; i++ {
		builder.WriteString(fmt.Sprintf("%v, ", a.arr[i]))
	}
	builder.WriteString(fmt.Sprintf("%v]", a.arr[i]))
	return builder.String()
}

func (a *Array[T]) Size() int {
	return len(a.arr)
}

func (a *Array[T]) Capacity() int {
	return cap(a.arr)
}

func (a *Array[T]) AddElems(elems ...T) {
	for _, elem := range elems {
		a.arr = append(a.arr, elem)
	}
}

func (a *Array[T]) At(pos int) (T, error) {
	var zero T
	if pos < 0 || pos > len(a.arr) {
		return zero, errors.New("index out of bound")
	}
	return a.arr[pos], nil
}

func (a *Array[T]) InsertElems(pos int, elems ...T) error {
	if pos < 0 || pos > len(a.arr) {
		return errors.New("index out of bound")
	}
	prev := make([]T, pos+len(elems))
	copy(prev, a.arr[:pos])
	for i, elem := range elems {
		prev[i+pos] = elem
	}
	a.arr = append(prev, a.arr[pos:]...)
	return nil
}

func (a *Array[T]) DeleteElems(pos int, count int) error {
	if pos < 0 || pos > len(a.arr) {
		return errors.New("index out of bound")
	}
	if count < 0 || count > len(a.arr[pos:]) {
		return errors.New("invalid count value")
	}
	a.arr = append(a.arr[:pos], a.arr[pos+count:]...)
	return nil
}
