package main

import (
	"fmt"
)

func main()  {
	var (
		a int
		b string
		c [3] float32

	)
	a = 1
	b = "卧槽了"
	c[1] = 2

	fmt.Println(c)
	fmt.Println(a)
	fmt.Println(b)

}
