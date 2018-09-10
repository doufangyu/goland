package main

import "fmt"

//自动推导变量类型
//简单计算
func main()  {
	var a  = 10
	var b  = 5
	var c float64 = 0.17
	var d  = float64(a-b) / c
	var e  = float64(a-b) * c
	fmt.Println(d)
	fmt.Println(e)
}
