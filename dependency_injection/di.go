package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	// Fprintf 构建一个 string 给 writer
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHandler))
}
