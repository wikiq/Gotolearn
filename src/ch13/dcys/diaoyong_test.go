package dcys

import (
	"fmt"
	"testing"
	"time"
)

func TestDiaoYong(t *testing.T) {
	switch time.Now().Weekday().String() {
	//time.Now().Weekday().String()是什么意思？
	//time.Now().Weekday().String()是获取当前时间的星期几的意思
	//首先获取日期，然后获取星期几，最后转换成字符串
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's time to learn some Go.")
	default:
		fmt.Println("It's the weekend, time to rest!")
	}

	fmt.Println(time.Now().Weekday().String())

}
