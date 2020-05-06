package main

import "fmt"

const (
    Big   = 1 << 100	//2^100
	Small = Big >> 99	//2^(100-99) = 2^1 = 2
	intMax = 1<<63 -1
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
    return x * 0.1
}

func main() {
	fmt.Println(Small)
	//fmt.Println(Big)	//overflows int error 발생
	fmt.Println(intMax)
    fmt.Println(needInt(Small))
    fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	//fmt.Println(needInt(Big))	//역시 overflows int error 발생
}
