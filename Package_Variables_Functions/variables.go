package main

import "fmt"

var x, y, z int
var c, python, java bool
var golang float32

var(
	a int
	b string
)

func main() {
	fmt.Println(x, y, z, c, python, java)
	fmt.Println(a,b)
	fmt.Scanln(&golang)
	fmt.Printf("You typed %f \n", golang)
}
