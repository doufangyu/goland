package main

import "fmt"

func main() {
	//a := 10
	//整型变量
	var p *int
	//指针变量
	//p = &a
	//*p = 100
	//等于操作 a = 100
	p = new(int)
	fmt.Printf("p = %v\n", p)
	*p = 6000
	fmt.Printf("*p = %v\n", *p)

	q := new(int)
	//自动推导类型
	*q = 8000
	fmt.Printf("*q = %v\n", *q)
}
