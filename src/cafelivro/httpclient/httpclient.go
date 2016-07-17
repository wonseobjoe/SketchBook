package httpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"net/url"
)

func Execute() {

	reqUrl := "https://openapi.naver.com/v1/search/news.xml"

	req, err := http.NewRequest("GET", reqUrl, nil)
	req.Header.Add("X-Naver-Client-Id", "fZJzwNBQ3KeDQ68wRgvv")
	req.Header.Add("X-Naver-Client-Secret", "S6fOmDjorn")

	query := req.URL.Query()
	query.Set("query", "이대호")
	query.Add("start", "1")
	query.Add("display", "10")
	query.Add("sort", "sim")
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error : %s", err)
	}

	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		// 결과 출력
		bytes, _ := ioutil.ReadAll(resp.Body)
		str := string(bytes) //바이트를 문자열로
		fmt.Println(str)
	}

}
