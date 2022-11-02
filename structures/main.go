package main

import (
	"data-structures/binary_tree"
	"fmt"
)

//func removeNthFromEnd(list *list.List[int], n int) *list.List[int] {
//	head := list.First
//	if head.Next == nil && n != 1 {
//		return nil
//	}
//	curr1 := head
//	for n >= 0 && curr1 != nil {
//		n--
//		curr1 = curr1.Next
//	}
//	if curr1 == nil && n == 0 {
//		head = head.Next
//		return list
//	}
//	if n != -1 {
//		return nil
//	}
//	prev := head
//	for curr1 != nil {
//		curr1 = curr1.Next
//		prev = prev.Next
//	}
//	next := prev.Next.Next
//	prev.Next = next
//	return list
//
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

	//for x := range tr.Iter() {
	//	fmt.Printf("%v ", x)
	//}

	//l := new(list.List[int])
	//l.AddBack(1)
	//l.AddBack(2)
	////l.AddBack(3)
	////l.AddBack(4)
	////l.AddBack(5)
	//
	//fmt.Println(l)
	//l = removeNthFromEnd(l, 2)
	//fmt.Println(l)

}
