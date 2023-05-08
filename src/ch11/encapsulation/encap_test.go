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

func (e *Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s///Name:%s///Age:%d", e.Id, e.Name, e.Age)
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	t.Log(e.String())
}

func (e Employee) String1() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s//Name:%s//Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	t.Log(e.String1())
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

/*--func解释--*/
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

//值传递和指针传递哪个性能更好？
//值传递更好，因为指针传递需要解引用，而值传递不需要解引用
//为什么指针传递需要解引用？
//因为指针传递是传递的地址，所以需要解引用
//为什么值传递不需要解引用？
//因为值传递是传递的值，所以不需要解引用
//哪种需要复制？
//值传递需要复制，指针传递不需要复制
//为什么值传递需要复制？
//因为值传递是传递的值，所以需要复制
