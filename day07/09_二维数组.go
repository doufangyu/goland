package main

import "fmt"

func main() {

	var a [3][4]int
	//有三个班级，每个班级有4个人
	// 1 - 1234
	//2 - 5678
	//3 - 9 10 11 12
	k := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			k++
			a[i][j] = k
			fmt.Printf("a[%d][%d] = %d, ", i, j, a[i][j])
		}
		fmt.Println()
	}
	fmt.Println("a = ", a)
	//另外一种写法
	//有三个元素，每个元素又是一维数组[4]int
	b := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println("b= ", b)

	//部分初始化，没有初始化的如果是int 全部补0

	c := [3][4]int{{1, 2, 3}, {5, 6}, {10}}
	fmt.Println("c= ", c)

	d := [3][4]int{{1, 2, 3, 4}}
	fmt.Println("d=", d)

	e := [3][4]int{1: {5, 6, 7, 8}}
	fmt.Println("e = ", e)
}
