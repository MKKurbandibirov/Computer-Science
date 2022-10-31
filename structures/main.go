package main

import (
	"data-structures/list"
	"fmt"
)

func main() {
	var l list.List[int]

	l.AddBack(12)
	fmt.Println(l.String())
	l.AddBack(32)
	l.AddBack(45)
	l.AddBack(56)
	for x := range l.Iter() {
		fmt.Println(x)
	}
	fmt.Println(l.String())

	l.DeleteElem(1)
	fmt.Println(l.String())

	fmt.Println(l.At(2))
}
