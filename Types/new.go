package main

import "fmt"

type Vertex struct {
    X, Y int
}

func main() {
	v := new(Vertex)
	var w *Vertex= new(Vertex)
    var z Vertex= Vertex{0,0}
	fmt.Println(v)	//&{0,0}
	fmt.Println(w)	//&{0,0}
    fmt.Println(z)	//{0,0}
    v.X, v.Y = 11, 9
	z.X,z.Y = 10, 8
    fmt.Println(v)	////&{0,0}
    fmt.Println(w)	////&{0,0}
	fmt.Println(z)	//{10,8}
}
