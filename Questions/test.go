package questions

import (
	"fmt"
	"testing"
)

func TestNumMovesStonesII(t *testing.T) {
	stones := []int{6, 5, 4, 3, 10}
	answ := NumMovesStonesII(stones)
	fmt.Println(answ)
}
