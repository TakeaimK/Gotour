package main

import (
    "fmt"
    "math"
)

func sqrt(x float64) string {
    if x < 0 {
        return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))

	var temp float64
	fmt.Scan(&temp)
	if temp < 0{
		fmt.Println(math.Sqrt(math.Abs(temp)),"i")
	}else{
		fmt.Println(math.Sqrt(math.Abs(temp)))
	}
}
