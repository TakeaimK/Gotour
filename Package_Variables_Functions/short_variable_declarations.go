package main

import "fmt"

func main() {
	var x, y, z int = 1, 2, 3
	w := 8	//함수 안에서만 사용 가능!
    c, python, java := true, false, "no!"
	s := "iPhone"

	fmt.Println(x, y, z, c, python, java)
	
	fmt.Printf("Apple %s %d", s, w)
}
