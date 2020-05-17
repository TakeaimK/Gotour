package main

import "fmt"

func out(ch chan int){
    fmt.Println(<-ch)   //1
}

func main() {
    c := make(chan int, 2)
    c <- 1
    c <- 2
    fmt.Println("Warning!")
  
    go out(c)
    c <- 3	
    
    fmt.Println(<-c)    //2
    fmt.Println(<-c)    //3
    fmt.Println("Warning!")
}