package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent", time.Since(start).Seconds())
		return ret
	}
}

// timeSpent是一个干嘛的函数？
// 给我解释一下“func timeSpent(inner func(op int) int) func(op int) int ”这段话
// timeSpent是一个函数，它的参数是一个函数，返回值也是一个函数
func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

//结果每次都不一样，因为rand.Intn()是随机数

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
		fmt.Println(ret)
	}
	return ret
}

func TestVarParam(t *testing.T) {

	t.Log(Sum(1, 2, 3, 4, 5))
	t.Log(Sum(1, 2, 3, 4))
}

func Clear() {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T) {
	defer Clear()
	/*defer func() {
		t.Log("Clear resources with func")
	}()*/
	t.Log("Started")
	panic("Fatal error") // defer仍然会执行
	fmt.Println("End")
}
