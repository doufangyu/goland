package main

import "fmt"

func main() {
	srcSlice := []int{1, 2}

	destSlice := []int{7, 7, 7, 7}

	//copy(destSlice,srcSlice)
	//fmt.Println("dest = ",destSlice)

	copy(srcSlice, destSlice)
	fmt.Println("src =", srcSlice)
}
