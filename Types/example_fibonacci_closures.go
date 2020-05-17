package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    x := 0
    y := 1
    start := true
    return func() int{
        if start == true{
            start = false
            return 1
        } else{
	        tmp := x+y
    	    x = y
        	y = tmp
        	return y
        }
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 15; i++ {
        fmt.Println(f())
    }
}