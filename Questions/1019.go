package questions

type ListNode struct {
	Val  int
	Next *ListNode
}

func NextLargerNodes(head *ListNode) []int {
	//reverse the list first
	newHead := reverseList(head)
	answ := []int{}
	increaseQueue := []int{}
	for newHead != nil {
		for len(increaseQueue) > 0 && newHead.Val >= increaseQueue[0] {
			increaseQueue = increaseQueue[1:]
		}
		if len(increaseQueue) > 0 {
			answ = append([]int{increaseQueue[0]}, answ...)
		} else {
			answ = append([]int{0}, answ...)
		}
		increaseQueue = append([]int{newHead.Val}, increaseQueue...)
		newHead = newHead.Next
	}
	return answ
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	for head != nil {
		next := head.Next
		(*head).Next = pre
		pre = head
		head = next
	}
	return pre
}
