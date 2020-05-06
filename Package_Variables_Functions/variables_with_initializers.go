package main

import "fmt"

var x, y, z int = 1, 2, 3
var c, python, java = true, false, "no!"

func main() {
	fmt.Println(x, y, z, c, python, java)
	
	var i = 100
	var s = "Kakao Enterprise"

	fmt.Println("나는",i,"% 확신을 가지고 있다.")
	fmt.Print(s)
}
