package dx01

import (
	"fmt"
	"testing"
)

func TestFuncOne(t *testing.T) {
	number1 := 50
	number2 := 80
	fmt.Println(sumMath(20, 30))
	fmt.Println(sumMath(number1, number2))
	intReturn(12, 12, 12, 12, 12)
}

func sumMath(a, b int) int {
	return a + b
}

func intReturn(args ...int) int {
	for i := 0; i < len(args); i++ {
		fmt.Println(i, args[i])
	}
	return 0
}

func TestAdressDone(t *testing.T) {
	var num int = 50
	fmt.Println(&num)
	//AdressTest(&num)
	fmt.Println(num)
}

func AdressTest(num int) {
	//*num = 30
	//num = 20
	fmt.Println(num)
}

func TestFuncThree(t *testing.T) {
	a := AdressTest
	fmt.Printf("a的类型是：%T,AdressTest的类型是：%T,TestAD对应的类型是：%T", a, AdressTest, TestAdressDone)
	a(10)
}
