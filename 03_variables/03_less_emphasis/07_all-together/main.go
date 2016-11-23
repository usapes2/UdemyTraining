package main

import "fmt"


var a = "this is stored in the variable a" //infered type , defined in package
var b,c string = " stored in b","stored in c" // defined in package
var d string // just defined

func main() {
	d ="stored in d"

	var e = 42
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(d)
	fmt.Println(e)
}