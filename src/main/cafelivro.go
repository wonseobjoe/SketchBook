package main

import (
<<<<<<< HEAD
	//"cafelivro/httpclient"
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
	//httpclient.Execute()
=======
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

>>>>>>> 25bdae1bef2acaf5545b61c22e3326e2cf25be93
}
