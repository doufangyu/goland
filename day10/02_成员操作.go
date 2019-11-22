package main

import "fmt"

type Person1 struct {
	name string
	sex  byte
	age  int
}

type Student1 struct {
	Person1 //这里类型没有字段，匿名字段 继承了person所有成员
	addr    string
	id      int
}

func main() {
	s1 := Student1{Person1{"mike", 'm', 18}, "beijing", 1}
	//对象成员的操作方式1
	s1.sex = 'f'
	s1.age = 222
	s1.name = "chaiio"
	s1.id = 7777

	//对象成员匿名字段造作方式二

	s1.Person1 = Person1{"tom", 'm', 19}
	fmt.Println(s1.Person1.name)
	fmt.Println(s1.name)
	fmt.Println(s1)
}
