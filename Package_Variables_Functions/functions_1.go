package main

import "fmt"

func add(x int, y int) int {
    return x + y
}

func add2(a int, b int) {
	fmt.Println(a+b)
}

func ret777 () int{
	return 777
}

func main() {
	fmt.Println(add(42, 13))	//55
	add2(6,4)					//10
	fmt.Println(ret777())		//777
	fmt.Println(ret777)			//0x49dd50
}
