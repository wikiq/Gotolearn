package _1Func

import (
	"fmt"
	"testing"
)

func swap(x, y string) (string, string) {
	return y, x
}

func TestMultiValue(t *testing.T) {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
