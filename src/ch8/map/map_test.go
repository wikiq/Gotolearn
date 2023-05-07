package map_ext

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
	//m[1](2)是什么意思？
	//m[1](2)是调用m[1]这个函数，参数是2
	//m[2](2)是什么意思？
	//m[2](2)是调用m[2]这个函数，参数是2
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	mySet[2] = true
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	t.Log(len(mySet))
	delete(mySet, 1)
	t.Log(len(mySet))
	t.Log(mySet)
	if mySet[1] {
		t.Logf("%d is existing", 1)
	} else {
		t.Logf("%d is not existing", 1)
	}
}
