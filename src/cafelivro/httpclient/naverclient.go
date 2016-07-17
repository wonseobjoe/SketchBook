package httpclient

import (
	"io/ioutil"
	"net/http"
)

const (
	NaverUrl = "https://google.search.com/v1/search/"
)

func Execute() {
	bytes, statuscode := GetNaverSearchResult("news", "이대호", "1", "10", "sim")

	println(statuscode)
	if statuscode == 200 {
		str := string(bytes) //바이트를 문자열로
		println(str)
	}

}

func GetNaverSearchResult(scope string, keyWord string, startNum string, displayNum string, sortType string) ([]byte, int) {

	reqUrl := NaverUrl + scope + ".xml"

	req, err := http.NewRequest("GET", reqUrl, nil)
	req.Header.Add("X-Naver-Client-Id", "fZJzwNBQ3KeDQ68wRgvv")
	req.Header.Add("X-Naver-Client-Secret", "S6fOmDjorn")

	query := req.URL.Query()
	query.Set("query", keyWord)
	query.Add("start", startNum)
	query.Add("display", displayNum)
	query.Add("sort", sortType)
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		println("Error : %s", err)
		return nil, -1
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		bytes, _ := ioutil.ReadAll(resp.Body)
		return bytes, resp.StatusCode
		//		str := string(bytes) //바이트를 문자열로
		//		fmt.Println(str)
	}
	return nil, resp.StatusCode
}
