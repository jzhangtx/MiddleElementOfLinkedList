package main

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

func GetArrayFromInput(count int, prompt string) []int {
	if count == 0 {
		return []int{}
	}

	fmt.Print(prompt)

	v := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&v[i])
	}

	return v
}

func GetListFromArray(v []int) *ListNode {
	if len(v) == 0 {
		return nil
	}

	if len(v) == 1 {
		return &ListNode{v[0], nil}
	}

	pHead := &ListNode{v[0], nil}
	pCur := pHead
	for i := 1; i < len(v); i++ {
		pCur.next = &ListNode{v[i], nil}
		pCur = pCur.next
	}

	return pHead
}

func CountLinkedList(list *ListNode) int {
	count := 0
	for ; list != nil; list = list.next {
		count++
	}

	return count
}

func GetNodeFromIndex(list *ListNode, index int) *ListNode {
	for i := 0; i < index; i++ {
		list = list.next
	}

	return list
}

func GetMiddleElementOfLinkedList(list *ListNode) int {
	count := CountLinkedList(list)
	index := (count - 1) / 2

	return GetNodeFromIndex(list, index).data
}

func main() {
	for {
		fmt.Print("Number of elements: ")
		var count int
		fmt.Scan(&count)
		if count == 0 {
			break
		}

		vec := GetArrayFromInput(count, "Please enter the list: ")
		pHead := GetListFromArray(vec)

		fmt.Print("The middle element of the linked list is ")
		fmt.Println(GetMiddleElementOfLinkedList(pHead))
	}
}
