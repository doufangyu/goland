package main

import "fmt"

func test(m map[int]string) {
	//map和我们的切片一样 属于引用类型
	delete(m, 1)
}

func main() {
	m := map[int]string{1: "jack", 2: "bbbb", 3: "qweqwe"}
	fmt.Println("m = ", m)

	test(m)
	fmt.Println(m)
}
