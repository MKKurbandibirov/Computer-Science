package stack

import (
	"errors"
	"fmt"
	"strings"
)

type Stack[T any] struct {
	arr []T
}

func (s *Stack[T]) Iter() <-chan T {
	iter := make(chan T)
	go func() {
		for i := 0; i < s.Size(); i++ {
			iter <- s.arr[i]
		}
		close(iter)
	}()
	return iter
}

func (s *Stack[T]) String() string {
	var builder strings.Builder
	var offset int
	for i := s.Size() - 1; i >= 0; i-- {
		tmp := fmt.Sprintln(s.arr[i])
		if len(tmp)/2 > offset {
			offset = len(tmp) / 2
		}
	}
	builder.WriteString(strings.Repeat("-", offset*2))
	builder.WriteString("\n")
	for i := s.Size() - 1; i >= 0; i-- {
		tmp := fmt.Sprintln(s.arr[i])
		for j := 0; j < offset-len(tmp)/2; j++ {
			tmp = " " + tmp
		}
		builder.WriteString(tmp)
	}
	builder.WriteString(strings.Repeat("-", offset*2))
	builder.WriteString("\n")
	return builder.String()
}

func (s *Stack[T]) Size() int {
	return len(s.arr)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) PushBack(elem T) {
	s.arr = append(s.arr, elem)
}

func (s *Stack[T]) Pop() {
	s.arr = s.arr[:s.Size()-1]
}

func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("stack is empty")
	}
	return s.arr[s.Size()-1], nil
}
