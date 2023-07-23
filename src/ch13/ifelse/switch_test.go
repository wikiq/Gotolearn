package ifelse

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSwitch(t *testing.T) {

	sec := time.Now().Unix()
	rand.Seed(sec)
	//i := 1
	i := rand.Int31n(10)
	//rand.Seed(sec)是什么意思？
	//第14行的rand.Seed(sec)是设置随机数种子的意思
	//i := rand.Int31n(2)是什么意思？
	//第15行的i := rand.Int31n(2)是生成0到1之间的随机数的意思
	fmt.Print("i = ", i)
	switch i {
	case 0:
		fmt.Print("zero...")
	case 1:
		fmt.Print("one...")
	case 2:
		fmt.Print("two...")
	default:
		fmt.Print("no match...")
	}
}
