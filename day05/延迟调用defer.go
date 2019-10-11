package main

import "fmt"

func main() {
	//defer 后进先出
	//先进后出
	defer fmt.Println("bbb")
	defer fmt.Println("ccc")
	fmt.Println("aaa ")
}
