package httpclient

import (
	"fmt"
	"net/http"
	//"net/url"
)

func Execute() {

	/*
		//aa := "https://openapi.naver.com/v1/search/news.xml?query=스페인&start=1&display=5"
		reqUrl := "https://openapi.naver.com/v1/search/news.xml"
		parseUrl, err := url.Parse(reqUrl)
		if err != nil {
			log.Fatal(err)
		}

		//	fmt.Println(parseUrl.Scheme)
		//	fmt.Println(parseUrl.Host)
		//	fmt.Println(parseUrl.Query())

		query := parseUrl.Query()
		query.Set("query", "김연아")
		query.Set("start", "1")
		query.Set("display", "5")
		query.Set("sort", "sim")

		parseUrl.RawQuery = query.Encode()
		fmt.Println(parseUrl)
	*/
	client := &http.Client{}

	reqUrl := "https://openapi.naver.com/v1/search/blog.xml?query=스페인&start=1&display=5"

	req, err := http.NewRequest("GET", reqUrl, nil)
	req.Header.Add("X-Naver-Client-Id", "fZJzwNBQ3KeDQ68wRgvv")
	req.Header.Add("X-Naver-Client-Secret", "S6fOmDjorn")

	fmt.Println(req.URL.RawQuery)
	query := req.URL.Query()

	req.URL.RequestURI()
	query.Set("query", "스페인")
	req.URL.RawQuery = query.Encode()

	fmt.Println(req.URL.RawQuery)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error : %s", err)
	}
	fmt.Println(resp)
	fmt.Println(resp.Status)

}
