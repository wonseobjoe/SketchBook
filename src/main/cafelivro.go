package main

import (
	//	"cafelivro/web"
	//	"net/http"
	//	"os"
	"cafelivro/httpclient"
)

func main() {
	httpclient.Execute()
	/*
		port := os.Getenv("PORT")
		//port = "80"

		http.Handle("/css/", new(web.StaticHandler))
		http.Handle("/fonts/", new(web.StaticHandler))
		http.Handle("/js/", new(web.StaticHandler))
		http.HandleFunc("/", web.IndexHandler)
		http.ListenAndServe(":"+port, nil)
	*/

}
