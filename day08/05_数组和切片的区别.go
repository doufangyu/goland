package main

import "fmt"

func main() {
	//切片和数组的区别

	a := [5]int{}
	//数组 数组[] 的长度是固定的常量，数组不能修改长度， len和 cap 固定
	fmt.Printf("len=%d, ca[=%d\n", len(a), cap(a))
	//切片是不固定长度的 不需要指定固定长度 可以增加
	//切片[]里边为空，[...]切片的长度不固定 可以追加
	s := []int{}
	//底层引入一个数组 不能装载下 就会新引入地址 不过对切片本身来说 不变
	//只要知道切片不计入长度就可以
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
	s = append(s, 1)
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
	s = append(s, 1)
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
	s = append(s, 1)
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
	s = append(s, 1)
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
	s = append(s, 1)
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
}
