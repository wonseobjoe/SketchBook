package main

import (
	//"sketchbook.org/web/example"
	"net/http"

	"sketchbook.org/web"
)

func main() {
	//example.Example1()
	//example.Example2()
	//example.Example3()
	http.Handle("/css/", new(web.StaticHandler))
	http.Handle("/fonts/", new(web.StaticHandler))
	http.Handle("/js/", new(web.StaticHandler))
	http.HandleFunc("/", web.IndexHandler)
	http.ListenAndServe(":8080", nil)
}
