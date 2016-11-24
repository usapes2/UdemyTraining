package main

import "fmt"

func main() {
	x := 12
	fmt.Println(x)

	{
		fmt.Println("this is inside mustashes ", x)
		y := "New variable inside mustashes"
		fmt.Println(y)
	}
	//fmt.Println(y) //outside scope of y
}
