package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func CryptoMd5(source string) string {
	h := md5.New()
	h.Write([]byte(source))
	return hex.EncodeToString(h.Sum(nil))
}

func CompareMd5(source, md5 string) bool {
	if CryptoMd5(source) == md5 {
		return true
	}
	return false
}
