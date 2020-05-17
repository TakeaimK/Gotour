package main

import "fmt"

func main() {
    m := make(map[string]int)

    m["Answer"] = 42
    fmt.Println("The value:", m["Answer"])
    //The value: 42

    m["Answer"] = 48
    fmt.Println("The value:", m["Answer"])
    //The value: 48

    delete(m, "Answer")
    fmt.Println("The value:", m["Answer"])
    //The value: 0

    v, ok := m["Answer"]
    fmt.Println("The value:", v, "Present?", ok)
    //The value: 0 Present? false

    m["korea"] = 82
    fmt.Println("The value:", m["korea"])
    //The value: 82

    w, okay := m["korea"]
    fmt.Println("The value:", w, "Present?", okay)
    //The value: 82 Present? true
}