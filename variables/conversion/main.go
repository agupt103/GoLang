package main

import (
	"fmt"
)
var a int
type hotdog int
var b hotdog 
func main() {
	a = 42
	b = 43
	fmt.Println("Value of a = ",a)
	fmt.Println("Value of b = ",b)
	a = int(b)
	fmt.Println("Value of a after conversion = ",a)
}