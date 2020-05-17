package main

import "fmt"

func main() {
    a := make([]int, 5)
    printSlice("a", a)  //a len=5 cap=5 [0 0 0 0 0]
    for i:=0; i<5; i++{
        a[i] = i
    }
    printSlice("a", a)  //a len=5 cap=5 [0 1 2 3 4]
    b := make([]int, 0, 5)
    printSlice("b", b)  //b len=0 cap=5 []
    /*
    for i:=0; i<5; i++{
        b[i] = i
    }
    */  //panic: runtime error: index out of range [0] with length 0
    c := b[:2]
    printSlice("c", c)  //c len=2 cap=5 [0 0]
    /*
    for i:=0; i<5; i++{
        c[i] = i
    }
    */  //panic: runtime error: index out of range [2] with length 2
    d := c[2:5]
    printSlice("d", d)  //d len=3 cap=3 [0 0 0]
}

func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n",
        s, len(x), cap(x), x)
}