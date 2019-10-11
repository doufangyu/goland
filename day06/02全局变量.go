package main

import "fmt"

var dou int = 100

func test3() {
	fmt.Println(dou)
}

func main() {
	fmt.Println(dou)
	test3()
}
