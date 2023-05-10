package _3named_return_value

import "testing"

func spiltvalue(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

func TestNamedReturnValue(t *testing.T) {
	t.Log(spiltvalue(17))
}
