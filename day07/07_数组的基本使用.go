package main

import "fmt"

func main() {
	//定义一个数组[10]int 和[5]int 是不同的类型
	//[数字]这个数字代表的数组的元素的个数
	var a [10]int
	//var b [5]int
	//fmt.Printf("len[a]=%d,len(b)=%d\n",len(a),len(b))
	//注意 定义数组 指定的数组的个数必须是常量
	//n := 10
	//const n int = 11
	//var c [n]int

	a[0] = 1
	i := 1
	a[i] = 10

	for i := 0; i < len(a); i++ {
		a[i] = i + 1
		fmt.Println(a[i])
	}

	//数组的操作 从0开始 最大长度为len()-1，不对称元素，这个数字叫下标

	for i, data := range a {
		fmt.Printf("a[%d]=%d\n", i, data)
	}
}
