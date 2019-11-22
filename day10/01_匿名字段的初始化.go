package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person //这里类型没有字段，匿名字段 继承了person所有成员
	addr   string
	id     int
}

func main() {
	//结构体里的顺序初始化 所有字段都需要初始化
	var s1 Student = Student{Person{"mike", 'M', 18}, "北京", 1}
	fmt.Println("s1 = ", s1)

	//自动类型推导
	s2 := Student{Person{"mike", 'm', 18}, "beijing", 2}
	fmt.Printf("s2 = %+v\n", s2)
	//指定类型初始化，没有初始化的常用类型 按照默认值进行赋值 ini 0 string 为空
	s3 := Student{id: 1}
	fmt.Printf("s3 = %+v\n", s3)

	s4 := Student{Person: Person{name: "张三"}, id: 1}
	fmt.Printf("s4 = %+v\n", s4)

	var s5 Student

	s5.name = "mike"
	s5.id = 1
	s5.age = 18
	s5.addr = "beijing"
	s5.sex = 'M'
	fmt.Printf("s5 = %+v\n", s5)

}
