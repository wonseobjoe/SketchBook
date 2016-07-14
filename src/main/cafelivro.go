package main

import (
	//"sketchbook.org/web/example"
	"net/http"
	"os"
	"sketchbook.org/web"
)

func main() {
	//example.Example1()
	//example.Example2()
	//example.Example3()
	port := os.Getenv("PORT")
	println("port : ", port)
	println("port : ", port)
	http.Handle("/css/", new(web.StaticHandler))
	http.Handle("/fonts/", new(web.StaticHandler))
	http.Handle("/js/", new(web.StaticHandler))
	http.HandleFunc("/", web.IndexHandler)
	http.ListenAndServe(":"+port, nil)
}
