package main

import (
	"algos-and-structures/structures/list"
	"fmt"
)

func main() {
	l := new(list.List[int])
	l.Insert(0, 1)
	l.Insert(0, 2)
	l.Insert(2, 3)
	fmt.Println(l)
}
