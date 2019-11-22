package main

import "fmt"

type Student2 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	//定义一个普通变量

	var s Student2
	//定义一个指针变量
	var p1 *Student2
	p1 = &s
	p1.age = 1
	fmt.Println(p1)

}
