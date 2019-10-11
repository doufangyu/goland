package main

import "fmt"

func main() {
	var a int = 10
	//每个变量有两层含义，第一个就是变量的内存，第二个是变量的地址
	fmt.Printf("a=%d\n", a)     //变量的内存
	fmt.Printf("&a = %v\n", &a) //变量的地址
	//保存地址必须要用指针（地址）
	var p *int
	//定一个变量P ，类型为int类型
	p = &a
	//指针指向谁，就把谁的地址赋值给谁
	fmt.Printf("P = %v\n,&a=%v\n,*p=%v\n", p, &a, *p)
	//fmt.Println(*p)
	*p = 1000
	fmt.Printf("P = %v\n,&a=%v\n,*p=%v\n", p, &a, *p)
}
