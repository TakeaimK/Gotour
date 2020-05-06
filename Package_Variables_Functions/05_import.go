package main

import (
    "fmt"
    "math"
)

func main() {
	fmt.Println("Now you have %g problems.", math.Nextafter(2, 3))
	//Nextafter(x,y) : x보다 y가 클 경우 표현할 수 있는 x보다 큰 실수
	fmt.Printf("20에 Root를 씌운 값은 %f 입니다.", math.Sqrt(20))
}
