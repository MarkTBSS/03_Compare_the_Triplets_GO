package main

import "fmt"

func compareTripletsPassByValue(a []int32, b []int32) []int32 {
	var pointOfEachSliceInScope []int32
	pointOfEachSliceInScope = []int32{0, 0}
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			pointOfEachSliceInScope[0]++
		}
		if b[i] > a[i] {
			pointOfEachSliceInScope[1]++
		}
	}
	return pointOfEachSliceInScope
}

/* func compareTripletsPassByReference(a []int32, b []int32, pointOfEachSlice *[]int32) {
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			(*pointOfEachSlice)[0]++
		}
		if b[i] > a[i] {
			(*pointOfEachSlice)[1]++
		}
	}
} */

func main() {
	var a []int32
	a = []int32{5, 6, 7}
	b := []int32{3, 6, 10}
	//fmt.Println(b)
	//สร้าตัวแปรให้ชนิดตรงกับที่ function return ออกมา
	var pointOfEachSlice []int32
	pointOfEachSlice = []int32{0, 0}
	//feed data เข้าไป
	pointOfEachSlice = compareTripletsPassByValue(a, b)
	//แสดงผลออกมาด้วย print
	fmt.Println(pointOfEachSlice)
}
