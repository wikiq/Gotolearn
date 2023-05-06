package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5}
	arr[1] = 7
	t.Log(arr[1], arr[2])
	t.Log(arr1[1], arr1[2])
	t.Log(arr3[1], arr3[2])
	t.Log(arr1, arr3)
	t.Log(len(arr1), len(arr3))
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	/*	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}*/
	for _, e := range arr3 {
		t.Log(e)
	}
	/*	for _, e := range arr3 {
		t.Log(e)
	}*/
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	arr3_sec := arr3[:3]
	t.Log(arr3_sec)
}



