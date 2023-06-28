package main

import (
	"log"
)

type ListNode struct {
     Val int
     Next *ListNode
}

func main() {

	l := ListNode{}
	l.Next = &ListNode{Next: nil}
	log.Println(hasCycle(&l))
}

func hasCycle(head *ListNode) bool {
    curr := head
    mapa := map[*ListNode]struct{}{}
	if curr == nil {
		return false
	}
    for {
        if curr.Next == nil {
            return false
        } else if _, ok := mapa[curr]; ok {
            return true
        }
		mapa[curr] = struct{}{}

        curr = curr.Next
    }
}