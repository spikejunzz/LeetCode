package questions

import (
	"fmt"
	"strconv"
)

func BaseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	s := string("")
	for n != 0 {
		r := n % -2
		if r < 0 {
			a := 1
			fmt.Println(a)
			r -= -2
		}
		n -= r
		s = strconv.Itoa(r) + s
		n /= -2
	}
	return s
}
