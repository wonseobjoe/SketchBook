<<<<<<< HEAD
package main

import (
	"cafelivro/web"
	"net/http"
	"os"
)

func main() {
	//example.Example1()
	//example.Example2()
	//example.Example3()
	port := os.Getenv("PORT")
	//port = "80"

	http.Handle("/css/", new(web.StaticHandler))
	http.Handle("/fonts/", new(web.StaticHandler))
	http.Handle("/js/", new(web.StaticHandler))
	http.HandleFunc("/", web.IndexHandler)
	http.ListenAndServe(":"+port, nil)

}
=======
package main

import (
	"cafelivro/web"
	"net/http"
	"os"
)

func main() {
	//example.Example1()
	//example.Example2()
	//example.Example3()
	port := os.Getenv("PORT")
	//port = "80"

	http.Handle("/css/", new(web.StaticHandler))
	http.Handle("/fonts/", new(web.StaticHandler))
	http.Handle("/js/", new(web.StaticHandler))
	http.HandleFunc("/", web.IndexHandler)
	http.ListenAndServe(":"+port, nil)

}
>>>>>>> 2b3e8bf30f00c1359cdb82cc02321e53c21373f3
