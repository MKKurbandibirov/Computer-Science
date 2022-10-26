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

func (a *Array[T]) AddElem(elem T) {
	a.arr = append(a.arr, elem)
}

func (a *Array[T]) AddElems(elems ...T) {
	for _, elem := range elems {
		a.arr = append(a.arr, elem)
	}
}

func (a *Array[T]) At(ind int) (T, error) {
	var zero T
	if ind < 0 || ind > len(a.arr) {
		return zero, errors.New("index out of bound")
	}
	return a.arr[ind], nil
}
