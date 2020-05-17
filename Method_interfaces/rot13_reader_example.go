package main

import (
    "io"
    "os"
    "strings"
    "fmt"
)

type rot13Reader struct {
    r io.Reader
}

func changeRot13(char byte)(cng byte){
    return
}


func (rot *rot13Reader)Read(data []byte) (len int, err error) {
    len, err = rot.r.Read(data)
    fmt.Println(data)
    // for i, char := range data{
    //     //data[i] = changeRot13(data[i])
    //     fmt.Println(data[i], char)
    // }
    return
}


func main() {
    s := strings.NewReader(
        "Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
