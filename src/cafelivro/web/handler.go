package web

import (
	//"encoding/json"
	"net/http"
)

func IndexHandler(writer http.ResponseWriter, req *http.Request) {
	indexContents, _ := loadIndexContents()
	renderTemplate(writer, "index", indexContents)
}

func NewsHandler(writer http.ResponseWriter, req *http.Request) {
	result := []byte(`{"name": "Zara Ali" ,"age" : "67" , "sex": "female"}`)

	writer.Write(result)

}
