package web

import (
	"cafelivro/httpclient"
	"io/ioutil"
	"net/http"
	"strconv"
)

func IndexHandler(writer http.ResponseWriter, req *http.Request) {
	indexContents, _ := loadIndexContents()
	renderTemplate(writer, "index", indexContents)
}

func SearchHandler(writer http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	scope := query.Get("scope")
	keyword := query.Get("keyword")
	pagenum, _ := strconv.Atoi(query.Get("pagenum"))
	startnum := (((pagenum - 1) * 10) + 1)
	sort := query.Get("sort")

	resp, _ := httpclient.GetNaverSearchResult(scope, keyword, strconv.Itoa(startnum), "10", sort)
	//resp, _ := httpclient.GetNaverSearchResult("news", "현재 날씨", "1", "10", "sim")

	if resp.StatusCode == 200 {
		bytes, _ := ioutil.ReadAll(resp.Body)
		writer.Write(bytes)
	}
}

func BasicHandler(writer http.ResponseWriter, req *http.Request) {
	result := []byte(`{"name": "Zara Ali" ,"age" : "67" , "sex": "female"}`)
	writer.Write(result)

}
