package crypto

import (
	"encoding/base64"
)

func Base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func Base64Decode(dec []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(dec))
}

func EncodeBase64(src string) string {
	if len(src) == 0 {
		return ""
	}
	return base64.StdEncoding.EncodeToString([]byte(src))
}

func DecodeBase64(dec string) (string, error) {
	if len(dec) == 0 {
		return "", nil
	}
	bytes, err := base64.StdEncoding.DecodeString(dec)
	return string(bytes), err
}
