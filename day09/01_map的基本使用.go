package main

import "fmt"

func main() {
	//定一个map变量 类型为mao[int]string

	var m1 map[int]string = make(map[int]string)

	fmt.Println("m1 = ", m1)
	m1[1] = "jake"

	fmt.Println("m1 = ", m1[1])

	m2 := make(map[int]string)
	m2[1] = "dou"

	fmt.Println(len(m2))
	fmt.Println("m2 = ", m2)

	///map 先给mao指定一个可以容纳长度 ，一旦炒股奥这个长度，从新分配底层空间

	m3 := make(map[int]string, 3)
	m3[1] = "dou"
	m3[2] = "fang"
	m3[3] = "yu"
	m3[4] = "chaiio"
	fmt.Println("m3 len=", len(m3))
	fmt.Println("m3 = ", m3)

}
