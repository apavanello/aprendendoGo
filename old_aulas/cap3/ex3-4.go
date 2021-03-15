package main

import (
	"fmt"
)

type car_value string

var x car_value

func main() {

	fmt.Println(x)
	fmt.Printf("x => %T \n", x)
	x = "42"
	fmt.Println(x)

}