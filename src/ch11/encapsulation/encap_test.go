package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployee0bj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	t.Log(e)
	e1 := Employee{Name: "Mike", Age: 30}
	t.Log(e1)
}

func (e *Employee) String() string {
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}

	t.Log(e.String())
}

//上面这是在给e赋值，下面这是在调用e的方法
//在给e赋值的时候，e是值传递，所以e的值不会改变
//e.String()是什么意思？
//调用e的String()方法
/*---*/
//func的格式是什么？
//func () {}
//func (参数列表) {}
//func (参数列表) 返回值列表 {}
//func (参数列表) (返回值列表) {}
//参数列表和返回值列表的格式是什么？
//参数列表：参数名 参数类型
//返回值列表：返回值名 返回值类型
//func (e Employee) String() string {}的意思是什么？
//e Employee是接收者，String()是方法名，string是返回值类型

/*func (e *Employee) String() string {
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}*/

//Employee和*Employee的区别是什么？
//Employee是值传递，*Employee是指针传递
//为什么要有这两种传递方式？
//因为有时候我们需要修改原来的值，有时候我们不需要修改原来的值
//比如我现在要修改原来的值，我应该用哪种传递方式？
//用指针传递
//比如我现在不需要修改原来的值，我应该用哪种传递方式？
//用值传递
