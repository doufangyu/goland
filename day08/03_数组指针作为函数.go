package main

import "fmt"

//*p = a 代表指针所指向大内存，就是实参
func modify1(p *[5]int) {
	(*p)[0] = 666
	fmt.Println("modify1 *p=", *p)
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	modify1(&a)
	fmt.Println("modufy1 a = ", a)

}
