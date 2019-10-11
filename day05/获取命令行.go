package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args //接收用户传递参数，都是以字符串传递
	n := len(list)
	fmt.Println(n)

	for i := 0; i < n; i++ {
		fmt.Printf("list[%d]=%s\n", i, list[i])
	}
	fmt.Println("______________------------________---__--_-_-__-_-")
	for i, data := range list {
		fmt.Printf("list[%d]=%s\n", i, data)
	}
}
