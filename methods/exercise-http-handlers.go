package main

import (
	"log"
	"net/http"
	"fmt"
)

// Str string
type Str string

func (s Str) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request){
		fmt.Fprintf(w, string(s))
}
// Text struct
type Text struct{
	Greeting string
	Punct string
	Who string
}

func (s Text) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request){
		fmt.Fprintf(w, s.Greeting + s.Punct + s.Who)
}

func main(){
	http.Handle("/string", Str("i`m a frayed knot."))
	http.Handle("/struct", &Text{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
	
}