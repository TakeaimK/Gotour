package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    m := make(map[string]int)
    arr := strings.Fields(s)
    
    for _, v:= range arr {
        
        _,t := m[v]
        if t == true {
            m[v]++
        } else{
            m[v] = 1
        }
    }
    
    return m
}

func main() {
    wc.Test(WordCount)
}
