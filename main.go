package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// listen requset by server
	mx := http.DefaultServeMux
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// send a picture as response
		// read response by ioutil
		// this file is in resources
		b, _ := ioutil.ReadFile("resources/hadi.bmp")
		w.Header().Add("content-type", "image/bmp")
		w.WriteHeader(http.StatusOk)
		w.Write(b)
	})
	// "/" means path in request (string)
	// the second parameter in inputs is a handler that gives a function

	// add new handle function
	mux.Handle("/books")

	http.ListenAndServe(":", nil)
	fmt.Println("Hi")
}

func handeleBooks(w http.ResponseWriter, r *http.Request) {
	// it use for handling request so it must receive a request(w) and
	// response writer and a pointer to request

	// various state in method
	switch r.Method {
	case "GET":

	case "POST":
		// if method isnt post and get , return an error for us
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
