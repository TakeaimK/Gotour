package main

import "fmt"

var x, y, z int
var c, python, java bool
var golang float32

func main() {
	fmt.Println(x, y, z, c, python, java)
	fmt.Scanln(&golang)
	fmt.Printf("You typed %f \n", golang)
}
