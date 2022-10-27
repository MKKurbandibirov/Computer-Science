package main

import (
	"data-structures/array"
	"fmt"
)

func main() {
	arr := new(array.Array[string])
	arr.AddElems("Haha")
	arr.AddElems("Hello")
	fmt.Println(arr)

	arri := new(array.Array[int])
	arri.AddElems(3)
	arri.AddElems(56)
	arri.AddElems(3, 4, 5, 6, 7, 8)
	//el, _ := arri.At(5)
	//fmt.Println(arri, el)
	fmt.Println(arri)

	arri.InsertElems(arri.Size(), 45, 56, 6, 66, 66)
	fmt.Println(arri)
	arri.DeleteElems(2, 4)
	fmt.Println(arri)
}
