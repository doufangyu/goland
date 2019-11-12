package main

import "fmt"

func main() {
	//自动推导类型 同时进行初始化
	s1 := []int{1, 2, 3, 4}
	fmt.Println("s1 =", s1)
	//添加到切片
	s1 = append(s1, 5)
	fmt.Println("s1 =", s1)

	//借助make的方式 创建切片 (类型 长度 容量)
	s2 := make([]int, 5, 10)
	//make(切片类型,切片长度，切片容量)
	fmt.Println("s2 =", s2)
	fmt.Printf("s2 len=%d,cap=%d\n", len(s2), cap(s2))

	//如果没有指定容量 那么容量和长度一样
	s3 := make([]int, 5)
	//make(切片类型,切片长度)
	fmt.Println("s3 =", s3)
	fmt.Printf("s3 len=%d,cap=%d\n", len(s3), cap(s3))
}
