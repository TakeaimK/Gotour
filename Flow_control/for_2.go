package main

import "fmt"

func main() {
	sum := 1
	rng := 5
    for sum < 1000 && rng > 0 {
		sum += sum
		rng--
    }
    fmt.Println(sum)
}
