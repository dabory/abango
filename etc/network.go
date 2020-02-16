package etc

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"strings"
	time "time"
)

func GetHttpResponse(method string, apiurl string, jsBytes []byte) ([]byte, []byte, error) {

	form := url.Values{}
	// form.Add("postvalues", string(kkk))
	// Values.Encode() encodes the values into "URL encoded" form sorted by key.
	// eForm := v.Encode()
	// fmt.Printf("v.Encode(): %v\n", s)
	reader := strings.NewReader(form.Encode())

	req, err := http.NewRequest(method, apiurl, reader)
	if err != nil {
		return nil, []byte("909"), MyErr("WERZDSVADFZ-http.NewRequest", err, true)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Endpoint-Agent", "abango-rest-api-v1.0")
	req.Header.Add("Accept-Language", "en-US")
	req.Header.Add("User-Agent", runtime.GOOS+"-"+runtime.Version()) // for checking OS Type in Server

	req.Body = ioutil.NopCloser(bytes.NewReader(jsBytes))

	// Client객체에서 Request 실행
	client := &http.Client{
		Timeout: time.Second * 20, //Otherwirse, it can cause crash without this line. Must Must.
	} // Normal is 10 but extend 20 on 1 Dec 2018

	// fmt.Println(reflect.TypeOf(respo))
	resp, err := client.Do(req)
	if err != nil {
		return nil, []byte("909"), MyErr("WERZDSVXBDCZSRE-client.Do "+apiurl, err, true)
	}
	defer resp.Body.Close()

	byteRtn, _ := ioutil.ReadAll(resp.Body)
	return byteRtn, []byte(strconv.Itoa(resp.StatusCode)), nil

}
