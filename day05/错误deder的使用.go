package main

import "fmt"

func test1(x int) {
	result := 1000 / x
	fmt.Println(result)
}

func main() {
	defer fmt.Println("aaaa")
	defer fmt.Println("bbbb")
	defer test1(0)
	defer fmt.Println("cccc")

}
