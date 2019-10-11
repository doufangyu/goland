package main

import "fmt"

const shuzuchangdu int = 5

func main() {
	//数组的长度必须是常量
	var id [shuzuchangdu]int
	//定义一个数组，同一类型的集合
	for i := 0; i < len(id); i++ {
		id[i] = i + 1
		fmt.Printf("id[%d]=%d\n", i, id[i])

	}
}
