package questions

import "sort"

func max(number1 int, number2 int) int {
	if number1 >= number2 {
		return number1
	} else {
		return number2
	}
}
func min(number1 int, number2 int) int {
	if number1 <= number2 {
		return number1
	} else {
		return number2
	}
}
func NumMovesStonesII(stones []int) []int {
	answ := make([]int, 2)
	sort.Ints(stones)
	stones_len := len(stones)
	//maximum steps;
	l1 := stones[stones_len-1] - stones[1] + 2 - stones_len
	l2 := stones[stones_len-2] - stones[0] + 2 - stones_len
	maximum := max(l1, l2)
	answ[1] = maximum
	//minimum steps;

	if l1 == 0 || l2 == 0 {
		answ[0] = min(2, answ[1])
		return answ
	}
	//sliding window
	maxCount := 0
	left := 0
	right := 0
	for right < stones_len {
		//cur_cnt := right - left + 1
		for stones[right]-stones[left]+1 > stones_len && right >= left {
			left++
		}
		maxCount = max(maxCount, right-left+1)
		right++
	}
	return []int{stones_len - maxCount, maximum}
}
