package main

import "fmt"

func main() {
	sum1 := 0
	sum2 := 0
    for i := 0; i < 10; i++ {
        sum1 += i
	}
	//sum2 = i*(i+1)/2	//i는 for문을 나오면 해제됨
	fmt.Println(sum1)
	fmt.Println(sum2)
}
