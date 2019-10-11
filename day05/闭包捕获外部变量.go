package main

import "fmt"

func main() {
	a := 10
	str := "make"

	func() {
		a = 666
		str = "go"
	}()
	fmt.Println(a, str)
}
