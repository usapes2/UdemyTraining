package main

import "fmt"

//var x = 0

func increment(x int) int {

	x++
	return x
}

func main() {
	var x =0
	fmt.Println(increment(x))
	fmt.Println(increment(x))
	fmt.Println(increment(x))
	fmt.Println(increment(x))
}
