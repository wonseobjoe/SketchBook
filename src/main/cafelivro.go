package main

import (
	//"cafelivro/httpclient"
	"cafelivro/web"
	"net/http"
	"os"
)

func main() {

	//httpclient.Execute()

	port := os.Getenv("PORT")
	//port = "8080"

	http.Handle("/css/", new(web.StaticHandler))
	http.Handle("/fonts/", new(web.StaticHandler))
	http.Handle("/js/", new(web.StaticHandler))
	http.Handle("/web/", new(web.StaticHandler))
	http.HandleFunc("/", web.IndexHandler)
	http.HandleFunc("/search/", web.SearchHandler)
	//http.HandleFunc("/contents/", web.ContentsHandler)
	http.ListenAndServe(":"+port, nil)

}
