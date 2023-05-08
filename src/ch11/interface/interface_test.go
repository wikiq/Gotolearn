package _interface

import "testing"

type programmer interface {
	WriteHelloWorld() string
}

//interface是什么？
//interface是一组method的组合，是duck type的体现
//interface是什么意思？
//interface是接口的意思
//interface的格式是什么？
//type 接口名 interface {
//	方法名1(参数列表) 返回值列表
//	方法名2(参数列表) 返回值列表
//	...
//}

type GoProgrammer struct {
}

//struct是什么？
//struct是一组数据的组合，是class的体现
//struct是什么意思？
//struct是结构体的意思
//struct的格式是什么？
//type 结构体名 struct {
//	字段名1 字段类型1
//	字段名2 字段类型2
//	...
//}

func (g *GoProgrammer) WriteHelloWorld() string {
	//return "fmt.Println(\"Hello World\")"
	return "Hello World"
}

func TestClient(t *testing.T) {
	var p programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
	//此处的p在GOProgrammer构造函数中是g，对吗？
	//对
}
