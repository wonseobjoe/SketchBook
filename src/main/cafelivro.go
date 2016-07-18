package main

import (
	"cafelivro/web"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	port = "80"

	http.Handle("/css/", new(web.StaticHandler))
	http.Handle("/fonts/", new(web.StaticHandler))
	http.Handle("/js/", new(web.StaticHandler))
	http.HandleFunc("/", web.IndexHandler)
	http.HandleFunc("/news/", web.NewsHandler)
	http.ListenAndServe(":"+port, nil)

}
