package main

import "fmt"

func test() func() int {
	var x int //x没有初始化 =0
	return func() int {
		x++
		return x * x
	}

}

func main() {
	//返回值是一个匿名函数，返回一个函数类型，通过f来调用返回的匿名函数，f来调用闭包函数
	//f不关心这些捕获的常量和变量是否超出作用域，这个变量x就会一直存在
	f := test()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	//重新初始化f函数
	f = test()
	fmt.Println(f())
}

//func test() int {
//	//函数调用的时候，x才会被分配空间，才初始化为0
//	var x int //X没有初始化，值为0
//	x ++
//	return x*x  	//函数调用完毕，x自动释放
//}
//
//func main()  {
//	fmt.Println(test())
//	fmt.Println(test())
//	fmt.Println(test())
//	fmt.Println(test())
//	//
//}
