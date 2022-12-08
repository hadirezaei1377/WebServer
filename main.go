package main

import (
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

	http.ListenAndServe(":", nil)
}
