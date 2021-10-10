package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//http handlers:

	//HandleFunc registered a function to a path on the default servemux..

	//It allows you to map a function to a path...
	// Defaultservemux ->

	//HandleFunc registers a handler for a given pattern in the defaultservemux..
	// It is a HTTP Handler, it does more than a

	//HTTP performant web server that has the ability to
	//Minimalist approach, uses GO Standard Library.

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		d, _ := ioutil.ReadAll(r.Body)

		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("goodbye")
	})

	http.ListenAndServe(":9090", nil)
}
