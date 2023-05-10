package _1Func

import (
	"fmt"
	"testing"
)

func add2(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add2(42, 13))
}

func TestFunctwo(t *testing.T) {
	fmt.Println(add2(42, 13))

}
