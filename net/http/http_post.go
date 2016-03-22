package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

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

func PostForm(url string, data url.Values) (response string, err error) {
	resp, err := http.PostForm(url, data)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func Post(url string, bodyType string, body io.Reader) (response string, err error) {
	resp, err := http.Post(url, bodyType, body)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(_body), nil
}
