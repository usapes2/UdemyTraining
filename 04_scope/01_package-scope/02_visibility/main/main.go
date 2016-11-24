package main

import (
	"fmt"
	"github.com/usapes2/UdemyTraining/04_scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	fmt.Println(vis.yourName)
	vis.PrintVar()

}
