package main

import "fmt"

func add(x int, y int) int {
    return x + y
}

func add_a_to_b(a int, b int) int{
	x:=b+1
	return b*x/2
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Print(add_a_to_b(1,10))
}
