package main

import "fmt"

type Person2 struct {
	name string
	sex  byte
	age  int
}

type Student2 struct {
	Person2 //这里类型没有字段，匿名字段 继承了person所有成员
	addr    string
	id      int
	name    string
}

func main() {
	//s1 := Student2{Person2{"mike",'m',18},"beijing",1,""}
	//声明(定义一个变量)
	var s Student2
	//就近原则
	s.name = "tom"
	//s.Person2.name = "jeck"

	fmt.Printf("s.name = %d,s.person.name = %d\n", s.name, s.Person2.name)

}
