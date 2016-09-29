package api

import (
	"testing"

	"encoding/json"
	"fmt"
	request "github.com/banerwai/global/bean"
	"github.com/banerwai/global/constant"
	"github.com/banerwai/gommon/crypto"
	"sort"
	"strconv"
	"time"
)

func TestCheckAPIJson(t *testing.T) {

	var APIRequest request.APIRequest

	var APIRequestHeader request.APIRequestHeader
	APIRequestHeader.Method = "invoke"
	APIRequestHeader.Ver = "1.0.0"
	APIRequest.Header = APIRequestHeader

	var APIRequestPayload request.APIRequestPayload
	APIRequestPayload.Iss = "Eric"
	APIRequestPayload.Iat = time.Now().Unix()
	APIRequestPayload.Aud = "www.banerwai.com"

	_mapParams := make(map[string]string)

	_mapParams["u"] = "ministor@126.com"
	_mapParams["p"] = "13811111111"
	APIRequestPayload.Params = _mapParams

	var _sortKeys []string
	for k := range _mapParams {
		_sortKeys = append(_sortKeys, k)
	}
	APIRequestPayload.SortKeys = _sortKeys
	APIRequest.Payload = APIRequestPayload

	var APIRequestSignature request.APIRequestSignature
	sort.Strings(_sortKeys)
	param := strconv.FormatInt(APIRequestPayload.Iat, 10) + "."
	for _, k := range _sortKeys {
		param += _mapParams[k]
	}
	param += constant.BanerwaiAPISignKey
	APIRequestSignature.Alg = "md5"
	APIRequestSignature.Sign = crypto.Md5(param)

	APIRequest.Signature = APIRequestSignature

	b, err := json.Marshal(APIRequest)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))

	fmt.Println(CheckAPIJson(string(b)))

	if !CheckAPIJson(string(b)) {
		t.Errorf("CheckAPIJson error")
	}
}
