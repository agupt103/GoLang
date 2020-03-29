package main

import (
	"fmt"
)

type hotdog int
var b hotdog 
func main() {
	b = 542
	fmt.Printf("%v\n", b)
	fmt.Printf("%T\n", b)
}