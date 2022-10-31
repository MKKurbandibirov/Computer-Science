package dequeue

import (
	"errors"
	"fmt"
	"strings"
)

type Dequeue[T any] struct {
	arr []T
}

func (d *Dequeue[T]) Iter() <-chan T {
	iter := make(chan T)
	go func() {
		for i := 0; i < d.Size(); i++ {
			iter <- d.arr[i]
		}
		close(iter)
	}()
	return iter
}

func (d *Dequeue[T]) String() string {
	var builder strings.Builder
	builder.WriteString("< ")
	for i := 0; i < d.Size()-1; i++ {
		builder.WriteString(fmt.Sprint(d.arr[i]))
		builder.WriteString(" - ")
	}
	builder.WriteString(fmt.Sprint(d.arr[d.Size()-1]))
	builder.WriteString(" >")
	return builder.String()
}

func (d *Dequeue[T]) Size() int {
	return len(d.arr)
}

func (d *Dequeue[T]) IsEmpty() bool {
	return d.Size() == 0
}

func (d *Dequeue[T]) AddFront(elem T) {
	d.arr = append([]T{elem}, d.arr...)
}

func (d *Dequeue[T]) AddBack(elem T) {
	d.arr = append(d.arr, elem)
}

func (d *Dequeue[T]) RemoveFront() error {
	if d.IsEmpty() {
		return errors.New("dequeue is empty")
	}
	d.arr = d.arr[1:]
	return nil
}

func (d *Dequeue[T]) RemoveBack() error {
	if d.IsEmpty() {
		return errors.New("dequeue is empty")
	}
	d.arr = d.arr[:d.Size()-1]
	return nil
}

func (d *Dequeue[T]) PeakFront() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, errors.New("dequeue is empty")
	}
	return d.arr[0], nil
}

func (d *Dequeue[T]) PeakBack() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, errors.New("dequeue is empty")
	}
	return d.arr[d.Size()-1], nil
}
