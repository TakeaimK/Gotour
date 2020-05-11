package main

import "fmt"

func main() {
    var z []int
    fmt.Println(z, len(z), cap(z))
    if z == nil {
        fmt.Println("nil!")
	}
	
	var a []int
	var b []int = make([]int, 0)
	var c []int = []int{}

	fmt.Println(a, b, c) // [] [] []

	fmt.Println(a == nil) // true
	fmt.Println(b == nil) // false
	fmt.Println(c == nil) // false
}
