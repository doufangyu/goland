package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InitData(s []int) {
	//设置种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}
}

func BubbleSort(s []int) {
	//冒泡排序
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func main() {
	//切片做参数是引用关系
	//引用类型  地址改变原内容值

	n := 10
	//创建一个切片
	s := make([]int, n, n)
	InitData(s)
	fmt.Println("排序前的值", s)
	BubbleSort(s)
	fmt.Println("排序后的值", s)
}
