package dx01

import (
	"fmt"
	"testing"
)

func TestIfOne(t *testing.T) {
	if count := 20; count <= 25 {
		fmt.Println(count, "是小于25的")
	}
}

func TestIfTwo(t *testing.T) {
	if count := 40; count <= 25 {
		fmt.Println(count, "是小于25的")
	} else {
		fmt.Println(count, "是大于25的")
	}
}

func TestIfThree(t *testing.T) {
	//count := 78
	if count := 65; count > 90 && count <= 100 {
		fmt.Println("A+")
	} else if count > 80 && count <= 90 {
		fmt.Println("B")
	} else if count > 60 && count <= 80 {
		fmt.Println("C")
	} else {
		fmt.Println("不及格")
	}
}
