package main

import "fmt"

const metersToYard float64 = 1.093

func main() {
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYard
	fmt.Println(meters, " meters is ", yards, " yards.")
}
