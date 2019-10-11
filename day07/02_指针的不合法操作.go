package main

import "fmt"

func main() {
	var p *int
	//指针变量
	p = nil
	fmt.Println("p = ", p)
	*p = 1000
	fmt.Println("*p = ", *p)
	//panic: runtime error: invalid memory address or nil pointer dereference
	//[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x10992da]
	//因为P没有任何指向

}
