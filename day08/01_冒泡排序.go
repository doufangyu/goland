package main

import (
	"fmt"
)

func main() {
	//冒泡排序 升序
	a := [8]int{24, 69, 80, 57, 13, 100, 1, 9}
	n := len(a)
	//for i := 0;i < n-1 ; i++ {
	//	for j := 0 ; j < n-1-i ; j++{
	//		if a[j]  > a[j+1]{
	//			a[j] , a[j+1] = a[j+1],a[j]
	//		}
	//	}
	//}

	//冒泡排序 降序
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j] < a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Printf("\n排序后：\n")
	for i := 0; i < n; i++ {
		fmt.Printf("%d\n", a[i])
	}
	fmt.Println()
}
