package main

import "fmt"

const Pi = 3.14

func main() {
    const World = "안녕"
    fmt.Println("Hello", World)
    fmt.Println("Happy", Pi, "Day")

    const Truth = true
    fmt.Println("Go rules?", Truth)

    test := "World!"
    fmt.Println("Hello", test)
    test = "세상!"
    fmt.Println("Hello", test)
    //World = "Hi"

}
