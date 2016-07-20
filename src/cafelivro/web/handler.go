package web

import (
	"cafelivro/httpclient"
	"io/ioutil"
	"net/http"
)

func IndexHandler(writer http.ResponseWriter, req *http.Request) {
	indexContents, _ := loadIndexContents()
	renderTemplate(writer, "index", indexContents)
}

func NewsHandler(writer http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	keyword := query.Get("keyword")
	println("keyword : ", keyword)
	resp, _ := httpclient.GetNaverSearchResult("news", keyword, "1", "10", "sim")

	if resp.StatusCode == 200 {
		bytes, _ := ioutil.ReadAll(resp.Body)
		writer.Write(bytes)
	}

}

func BasicHandler(writer http.ResponseWriter, req *http.Request) {
	result := []byte(`{"name": "Zara Ali" ,"age" : "67" , "sex": "female"}`)
	writer.Write(result)

}
