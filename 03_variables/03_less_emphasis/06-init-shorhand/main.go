package main

import "fmt"

func main() {
		// note : it is only can be done inside of a func
	message := "hello World"
	a, b, c := 1, false, 3
	d := 4
	e := true

	fmt.Printf("%T,%T,%T,%T,%T,%T\n", message, a, b, c, d, e)

}