package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// var HttpClient *http.Client

// func init() {
// 	HttpClient = http.DefaultClient
// }

func PostRawJson(finalURL string, req []byte, response interface{}) (err error) {
	httpResp, err := http.DefaultClient.Post(finalURL, "application/json; charset=utf-8", bytes.NewReader(req))
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	if err = json.NewDecoder(httpResp.Body).Decode(response); err != nil {
		return
	}
	return
}

func PostJsonObj(finalURL, _json string, response interface{}) (err error) {

	httpResp, err := http.DefaultClient.Post(finalURL, "application/json; charset=utf-8", bytes.NewReader([]byte(_json)))
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	if err = json.NewDecoder(httpResp.Body).Decode(response); err != nil {
		return
	}
	return
}

func PostJsonString(finalURL, _json string) (response string, err error) {

	httpResp, err := http.DefaultClient.Post(finalURL, "application/json; charset=utf-8", bytes.NewReader([]byte(_json)))
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	_responseBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return "", err
	}

	return string(_responseBody), nil
}
