package main

import "fmt"

func main() {
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//[low:hight:max] s[1:2:4] 取下标 从low开始
	//len = hight-low
	//cap = max-low
	s1 := array[:]
	//取出所有
	// [0:len(array):len(array)]
	//容量和长度是一致的
	fmt.Println("s1 = ", s1)
	fmt.Printf("len=%d,cap=%d\n", len(s1), cap(s1))

	//操作某个元素 和数组的操作方法一样
	date := array[1]
	fmt.Println("data = ", date)
	date1 := array[3:6:7]
	//取出 a[3] a[4] a[5]
	//len = 7-3 cap = 7-4
	fmt.Println("date1 = ", date1)
	fmt.Printf("len=%d,cap=%d\n", len(date1), cap(date1))

	date2 := array[:6]
	//从0开始到第6个
	fmt.Println("date2 =", date2)
	fmt.Printf("len=%d,cap=%d \n", len(date2), cap(date2))

	s4 := array[3:]
	fmt.Println("s4 =", s4)
	fmt.Printf("len=%d,cap=%d \n", len(s4), cap(s4))
}
