package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
    return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, ErrNegativeSqrt(f)
    }
    return math.Sqrt(f), nil
}

func main() {
    //fmt.Println(Sqrt(2))
    //fmt.Println(Sqrt(-2))
    if ans, err := Sqrt(2); err!=nil{
        fmt.Println(err)
    } else{
        fmt.Println(ans)
    }
    if ans, err := Sqrt(-2); err!=nil{
        fmt.Println(err)
    } else{
        fmt.Println(ans)
    }
}