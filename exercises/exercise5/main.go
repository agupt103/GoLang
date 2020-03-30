package main

import "fmt"

type test int
var x test
var y int

func main() {

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n",y)
}