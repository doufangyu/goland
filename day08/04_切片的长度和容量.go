package main

import "fmt"

func main() {
	//silce
	a := [5]int{1, 2, 3, 4, 5}
	//左包涵 右不包含
	s := a[0:3:5]
	//切片长度 = 3-0=3
	//容量 = 5-0=5
	fmt.Println("s=", s)
	fmt.Println("len(s)=", len(s))
	fmt.Println("cap(s)=", cap(s))
	l := len(a)
	s1 := a[1:3:l]
	fmt.Println("s1=", s1)
	fmt.Println("len(s1)=", len(s1))
	fmt.Println("cap(s1)=", cap(s1))
	fmt.Printf("%p", &s1)
	s1 = append(s1, 1)
	fmt.Println("s1=", s1)
	fmt.Println("len(s1)=", len(s1))
	fmt.Println("cap(s1)=", cap(s1))
	fmt.Printf("%p", &s1)
}
