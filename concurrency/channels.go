package main

import (
    "fmt"
    //"time"
    )

func sum(a []int, c chan int) {
    sum := 0
    for _, v := range a {
        sum += v
    }
    c <- sum // send sum to c
}

func main() {
    a := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int, 3)
    go sum(a[:len(a)/2], c)	//17
    go sum(a[len(a)/2:], c)	//-5
    go sum(a[:], c)			//12
    //time.Sleep(100 * time.Millisecond)
    x, y, z := <-c, <-c, <-c// receive from c

    fmt.Println(x, y, z)
}
