package main

import (
	"fmt"
)

type car_value int

var (
	x car_value 
	y int
)


func main() {

	fmt.Println(x)
	fmt.Printf("x => %T \n", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Printf("y => %v , %T ", y, y)


}