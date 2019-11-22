package main

import "fmt"

//定义一个结构体
type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	////顺序初始化，每个成员必须初始化
	//var s1 Student = Student{1,"mike",'M',18,"北京"}
	//fmt.Println("s1 =" ,s1)
	////指定成员初始化，没有初始化的成员 ，自动赋值为0
	//s2 := Student{name:"mike",addr:"北京"}
	//fmt.Println("s2 = ",s2)

	//定义一个结构体的普通变量
	var s Student
	//操作成员变量 需要用.进行操作
	s.id = 1
	s.name = "mike"
	s.sex = 'M'
	s.addr = "北京"
	s.age = 18

	fmt.Println("s = ", s)

}
