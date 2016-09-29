package api

import (
	"encoding/json"
	request "github.com/banerwai/global/bean"
	"github.com/banerwai/global/constant"
	"github.com/banerwai/gommon/crypto"
	"sort"
	"strconv"
)

// CheckAPIJson check api signature
func CheckAPIJson(jsonRequest string) bool {
	var request request.APIRequest
	json.Unmarshal([]byte(jsonRequest), &request)
	return CheckAPISignature(request)
}

// CheckAPISignature check api signature
func CheckAPISignature(request request.APIRequest) bool {
	var _sortKeys []string
	for _, v := range request.Payload.SortKeys {
		_sortKeys = append(_sortKeys, v)
	}
	sort.Strings(_sortKeys)
	param := strconv.FormatInt(request.Payload.Iat, 10) + "."
	for _, k := range _sortKeys {
		param += request.Payload.Params[k]
	}
	param += constant.BanerwaiAPISignKey
	return crypto.CompareMd5(param, request.Signature.Sign)
}
