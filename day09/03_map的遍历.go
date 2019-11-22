package main

import "fmt"

func main() {
	m1 := map[int]string{1: "make", 2: "jack", 3: "tom"}
	//第一个返回值为key，第二个返回值为value 遍历结果是无序的
	for key, value := range m1 {
		fmt.Printf("%d =====%s\n", key, value)
	}

	//如何判断一个key是否存在
	//第一个返回值 为key的所对应的value,第二个返回值为key是否，如果ok为true
	//value,ok := m1[1]
	if value, ok := m1[1]; ok == true {
		fmt.Println("m[1]=", value)
	} else {
		fmt.Println("key 不存在")
	}

	delete(m1, 2)
	//删除key为2的内容
	fmt.Println(m1)

}
