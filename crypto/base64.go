package crypto

import (
	"encoding/base64"
)

func Base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func Base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func EncodeBase64(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

func DecodeBase64(src string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(src)
	return string(bytes), err
}
