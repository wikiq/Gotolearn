package dx01

import (
	"fmt"
	"testing"
)

func TestSwitchOne(*testing.T) {
	score := 12
	switch score / 10 {
	case 10:
		fmt.Println("A")
	case 9:
		fmt.Println("B")
	case 8:
		fmt.Println("C")
	case 7:
		fmt.Println("D")
	default:
		fmt.Println("E, You should go home")
	}

}

func TestSwitchTwo(*testing.T) {
	sum := 0
	for i := 0; i < 101; i += 2 {
		sum += i
	}
	fmt.Println(sum)
}
