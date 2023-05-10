package _1Func

import (
	"fmt"
	"testing"
)

func add(x int, y int) int {
	return x + y
}

func TestFunc(t *testing.T) {
	fmt.Println(add(42, 13))
}
