package main

import (
	// "fmt"
	"net/http"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	// create a new div element
	div := js.Global.Get("document").Call("createElement", "div")

	// set the innerHTML of the div
	div.Set("innerHTML", "<h1>Hello, world!</h1>")

	// /*document.write("Hello world!");*/
	// js.Global.Get("document").Call("write", "Hello world!")

	// append the div to the body of the page
	js.Global.Get("document").Get("body").Call("appendChild", div)

	// start the server
	http.ListenAndServe(":8080", nil)
}
