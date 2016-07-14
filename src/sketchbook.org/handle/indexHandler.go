package handle

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, req *http.Request, title string) {
	println("aaaaaaaaaaaaaaaaaaaa")
	err := templ.ExecuteTemplate(w, "index.html", nil)
	println("err : ", err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

var templ = template.Must(template.ParseFiles("index.html"))
