package httpclient

import (
	//xj "github.com/basgys/goxml2json"
	"io/ioutil"
	"net/http"
)

const (
	NaverUrl          = "https://openapi.naver.com/v1/search/"
	NaverClientId     = "fZJzwNBQ3KeDQ68wRgvv"
	NaverCleintSecret = "S6fOmDjorn"
)

func Execute() {
	println("execute")
	xml, statuscode := GetNaverSearchResultXml("news", "이대호", "1", "2", "sim")

	println(statuscode)
	println("xml", xml)

	// bug

	//	json, statuscode2 := GetNaverSearchResultJson("news", "이대호", "1", "2", "sim")
	//
	//	println(statuscode2)
	//	println("json", json)

}

func GetNaverSearchResultXml(scope string, keyWord string, startNum string, displayNum string, sortType string) (string, int) {
	resp, err := GetNaverSearchResult(scope, keyWord, startNum, displayNum, sortType)

	if err != nil {
		println("Error : %s", err)
	}

	if resp.StatusCode == 200 {
		bytes, _ := ioutil.ReadAll(resp.Body)
		str := string(bytes) //바이트를 문자열로
		return str, resp.StatusCode
	}
	return "", resp.StatusCode
}

/*
func GetNaverSearchResultJson(scope string, keyWord string, startNum string, displayNum string, sortType string) (string, int) {
	resp, err := GetNaverSearchResult(scope, keyWord, startNum, displayNum, sortType)

	if err != nil {
		println("Error : %s", err)
	}

	if resp.StatusCode == 200 {
		json, _ := xj.Convert(resp.Body)
		return json.String(), resp.StatusCode
	}
	return "", resp.StatusCode

}
*/

func GetNaverSearchResult(scope string, keyWord string, startNum string, displayNum string, sortType string) (resp *http.Response, err error) {

	reqUrl := NaverUrl + scope + ".xml"
	req, err := http.NewRequest("GET", reqUrl, nil)
	req.Header.Add("X-Naver-Client-Id", NaverClientId)
	req.Header.Add("X-Naver-Client-Secret", NaverCleintSecret)

	query := req.URL.Query()
	query.Set("query", keyWord)
	query.Add("start", startNum)
	query.Add("display", displayNum)
	query.Add("sort", sortType)
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}
	return client.Do(req)
	//resp, err := client.Do(req)

	//	if err != nil {
	//		println("Error : %s", err)
	//		return nil, err
	//	}
	//
	//
	//
	//	defer resp.Body.Close()
	//
	//	if resp.StatusCode == 200 {
	//		ioutil.
	//		bytes, _ := ioutil.ReadAll(resp.Body)
	//		str := string(bytes) //바이트를 문자열로
	//		println("aaaaaaaaaaaaaaaaa", str)
	//		//return str, statuscode
	//	}
	//	return &resp.Body, resp.StatusCode
}
