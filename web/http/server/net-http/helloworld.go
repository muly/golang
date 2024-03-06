// basic hello world web application just using net/http package
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/welcome", welcomeHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		// dfdfd
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)

	w.Write([]byte("Hello World!"))
}
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome World!"))
}
