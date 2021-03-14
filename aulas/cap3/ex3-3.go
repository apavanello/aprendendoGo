package main

import (
	"fmt"
)

var (
	x = 42
	y = "James Bond"
	z = true
)

func main(){

	s := fmt.Sprintf(" x => %v \n y => %v \n z => %v" , x, y, z)
	fmt.Println(s)
}