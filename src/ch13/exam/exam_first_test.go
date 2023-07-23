package exam

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestExam(t *testing.T) {
	sum := sum(os.Args[1], os.Args[2])
	fmt.Println("Sum:", sum)
}

func sum(number1 string, number2 string) int {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	return int1 + int2
}

//上面的strconv.Atoi()函数是什么意思？
//strconv.Atoi()函数是将字符串转换成整数的函数
//strconv.Atoi()函数的格式是什么？
//strconv.Atoi()函数的格式是strconv.Atoi(字符串)
