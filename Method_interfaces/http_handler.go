package main

import (
    "fmt"
    "net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

type k struct{

}

func (s String) ServeHTTP(w http.ResponseWriter,r *http.Request) {
    fmt.Fprint(w, s)
}

func (t *Struct) ServeHTTP(w http.ResponseWriter,r *http.Request) {
    fmt.Fprint(w, t.Greeting, t.Punct, t.Who)
}

func (str k) ServeHTTP(w http.ResponseWriter,r *http.Request) {
    fmt.Fprint(w, "Hello, World!")
}

func main() {
    // your http.Handle calls here
    http.Handle("/", k{})
    http.Handle("/string", String("I'm a frayed knot."))
    http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
    http.ListenAndServe("localhost:4000", nil)
}
