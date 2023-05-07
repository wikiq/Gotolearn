package my_map

import (
	"testing"
)

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[0])
	t.Logf("len m1=%d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
	//为什么len(m3)是0？
	//因为make的时候，只是给了一个容量，但是没有给任何元素
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existing")
	}
	//ok:=m1[3]的结果是什么？
	//ok是false
	//为什么？
	//因为m1[3]不存在，所以ok是false
	//m1[3]的值是什么？
	//m1[3]的值是0
	//为什么？
	//因为m1[3]不存在，所以是0
}

func TestTravelMap(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	for k, v := range m {
		t.Log(k, v)
	}
	m2 := map[int]string{1: "one", 2: "two", 3: "three"}
	for k, v := range m2 {
		t.Log(k, v)
	}

	//k是每次加一，v是每次加一，对吗？
	//不对
	//为什么？
	//因为map是无序的，所以key和value的顺序是无所谓的

	//我在赋值的时候，为什么是m := map[string]int{"one": 1, "two": 2, "three": 3}，为什么“one”可以在前面
	//因为map是无序的，所以key和value的顺序是无所谓的
	//我在遍历的时候，为什么是for k, v := range m，为什么k在前面
	//因为map是无序的，所以key和value的顺序是无所谓的
	//所以无论我写“one”：1还是1：“one”，都是一样的，对吗？
	//对
	//上面的k对应的是key，v对应的是value，对吗？
	//对
	//为什么我现在写的是m2 := map[int]int{1: "one", 2: "two", 3: "three"}，为什么是错误的
	//因为key和value的类型不对

	//给我解释一下for k, v := range m2 {}是什么含义
	//遍历m2，k是key，v是value
	//这个循环的过程是
	//第一次循环，k=1，v="one"
	//第二次循环，k=2，v="two"
	//第三次循环，k=3，v="three"
	//所以我在循环的时候，k和v的顺序是无所谓的，对吗？
	//对
}

//给我写一个range的例子
//for k, v := range m2 {
//	t.Log(k, v)
//}
//不要是关于map的，就最基础的
//for i, v := range []int{1, 2, 3, 4} {
//	t.Log(i, v)
//}
//再基础一点，就单纯的数。
//for i := 0; i < 5; i++ {
//	t.Log(i)
//}
//这并没有用到range，写一个带有range的
//for i := range []int{1, 2, 3, 4} {
//	t.Log(i)
//}

func TestTracelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}

	for i, k := range []int{3, 3, 7, 4} {
		t.Log(i, k)
	}
}
