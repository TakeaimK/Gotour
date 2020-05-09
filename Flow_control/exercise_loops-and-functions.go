package main

import (
    "fmt"
    "math"
)

func Sqrt1(x float64, r int) float64 {
    z := float64(1)

    for i:=0; i<r; i++{
        z = z-(z*z-x)/(2*z)
    }
    return z
}

func Sqrt2(x float64) float64{
    z := float64(1)
 
    for {
        result := z-(z*z-x)/(2*z)

        if float64(result) == float64(z) {
                return result
        }
        z = result
    }
}

func main() {
    var num float64
    var rng int
    fmt.Print("Sqrt number : ")
    fmt.Scan(&num)
    fmt.Print("range : ")
    fmt.Scan(&rng)
    fmt.Println(Sqrt1(num, rng))
    fmt.Println(Sqrt2(num))
    fmt.Println(math.Sqrt(num))
}
