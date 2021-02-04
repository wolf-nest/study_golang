package ch06_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 2, 3, 4, 5}

	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}

	// 数组基本遍历
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	// 数组遍历 类似 foreach
	for idx, e := range arr3 {
		t.Log(idx, e)
	}
}

func TestArraySection(t *testing.T) {
	// 数组快速截取
	// a[开始索引(包含),结束索引(不包含)]
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3_sec := arr3[:3]  //截取从0-3的元素
	arr3_sec1 := arr3[3:] //截取从3开始到结束的 元素
	t.Log(arr3_sec, arr3_sec1)
}
