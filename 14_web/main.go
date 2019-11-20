// Go is great to build REST APIs and microservices

package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//can imiplement html inside
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}
func about(w http.ResponseWriter, r *http.Request) {
	//can imiplement html inside
	fmt.Fprintf(w, "<h1>About</h1>")
}

func main() {
	// kind of like a router and takes in a route and you can pass in a function to deal with that route
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server Starting...")

// to make it run:
	//takes in a port and second param is nil
http.ListenAndServe(":3000", nil)
}
