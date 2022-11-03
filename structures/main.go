package main

import (
	"data-structures/binary_tree"
	"fmt"
)

//func removeElements(list *list.List[int], val int) *list.List[int] {
//	head := list.First
//	for head.Value == val && head != nil {
//		head = head.Next
//	}
//	if head == nil {
//		return nil
//	}
//	prev := head
//	curr := head.Next
//	for curr != nil {
//		if curr.Value == val {
//			prev.Next = curr.Next
//			curr = curr.Next
//		}
//		prev = prev.Next
//		if curr != nil {
//			curr = curr.Next
//		}
//	}
//	return list
//}

func main() {
	tr := new(binary_tree.BinaryTree[int])
	tr.Insert(23, func(a, b int) bool {
		return a < b
	})
	tr.Insert(14, func(a, b int) bool {
		return a < b
	})
	tr.Insert(45, func(a, b int) bool {
		return a < b
	})
	tr.Insert(12, func(a, b int) bool {
		return a < b
	})
	tr.Insert(16, func(a, b int) bool {
		return a < b
	})

	tr.Insert(34, func(a, b int) bool {
		return a < b
	})
	tr.Insert(35, func(a, b int) bool {
		return a < b
	})
	tr.Insert(33, func(a, b int) bool {
		return a < b
	})
	tr.Insert(56, func(a, b int) bool {
		return a < b
	})

	fmt.Print(tr.String())

	tr.Delete(tr.Root.Left.Right)

	fmt.Print(tr.String())

	//for x := range tr.Iter() {
	//	fmt.Printf("%v ", x)
	//}

	//l := new(list.List[int])
	//l.AddBack(1)
	//l.AddBack(2)
	//l.AddBack(6)
	//l.AddBack(3)
	//l.AddBack(4)
	//l.AddBack(5)
	//l.AddBack(6)
	////
	//fmt.Println(l)
	//l = removeElements(l, 6)
	//fmt.Println(l)

}
