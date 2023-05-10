package _4variable

import "testing"

var i, j int = 1, 2

func TestVariablebegin(t *testing.T) {
	var c, python, java = true, false, "no!"
	k := 3
	t.Log(i, j, k, c, python, java)
}
