package main

import "fmt"

func main() {
	a := 10
	b := 20

	defer func(a, b int) {
		fmt.Printf("a=%d,b=%d", a, b)
	}(a, b)
	a = 111
	b = 222

	fmt.Printf("外部a=%d,b=%d", a, b)
}
