package main

import (
	"encoding/json"
	"fmt"
	"io"
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

// store a list of book as a handler
// define handeleBooks function and then struct
// finally a slice of books by var

type book struct {
	Isbn string `"json:isbn`
	Name string `"json:name`
}

var bList []book = []book{}

func handeleBooks(w http.ResponseWriter, r *http.Request) {
	// it use for handling request so it must receive a request(w) and
	// response writer and a pointer to request

	// various state in method
	switch r.Method {
	case "GET":

		// in this method we want to return a list of books(book list)
		// this action needs write method
		// change write method inputs to slice of books by json.marshal
		b, _ := json.Marshal(bList)
		w.Write(b)
		return

	case "POST":
		// post usually use for generate a resource
		// so we need body
		// what is the content in body request
		// make this body to array of bites
		newBook := book{}
		b, _ := io.ReadAll(r.Body)
		json.Unmarshal(b, &newBook)
		// we define a varibale named newbook in tpe book
		// and use its address

	default:
		// if method isnt post and get , return an error for us
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
