package main 

import (
	"fmt"
)

var (
	x int
	y string
	z bool
)

func main() {

	fmt.Printf(" x: %v \n y: %v \n z: %v \n", x, y, z)
	fmt.Printf(" x: %T \n y: %T \n z: %T \n", x, y, z) 

}