package main

import "fmt"

func test2(a int) {
	a = 100
	fmt.Println(a)
}

func main() {
	a := 100
	fmt.Println(a)
	{
		b := 20
		fmt.Println(b)
	}
	test2(3000)
}
