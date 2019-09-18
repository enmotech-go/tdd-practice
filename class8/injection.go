package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(buffer io.Writer, name string) {
	fmt.Fprintf(buffer, "Hello ,%s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe("127.0.0.1:5000", http.HandlerFunc(MyGreeterHandler))
}
