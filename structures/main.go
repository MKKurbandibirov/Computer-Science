package main

import (
	"data-structures/array"
	"fmt"
)

func main() {
	arr := new(array.Array[string])
	arr.AddElem("Haha")
	arr.AddElem("Hello")
	fmt.Println(arr)

	arri := new(array.Array[int])
	arri.AddElem(3)
	arri.AddElem(56)
	arri.AddElems(3, 4, 5, 6, 7, 8)
	el, _ := arri.At(5)
	fmt.Println(arri, el)
}
