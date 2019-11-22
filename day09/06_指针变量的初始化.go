package main

import "fmt"

type Person struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	//p1 存放的是地址 Person
	var p1 *Person = &Person{1, "mike", 'M', 18, "北京"}
	fmt.Println("p1=", p1)
	fmt.Printf("p1 =%p\n", p1)
	fmt.Printf("p1 =%v\n", p1)

	p2 := &Person{name: "mike", addr: "北京"}
	fmt.Printf("p2 type is %T\n", p2)
	fmt.Println("p2 = ", p2)

}
