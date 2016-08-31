package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(source string) string {
	h := md5.New()
	h.Write([]byte(source))
	return hex.EncodeToString(h.Sum(nil))
}

func CompareMd5(source, md5 string) bool {
	if Md5(source) == md5 {
		return true
	}
	return false
}

func DoubleMd5(source string) string {
	return Md5(Md5(source))
}

func CompareDoubleMd5(source, md5 string) bool {
	_temp := Md5(source)
	if Md5(_temp) == md5 {
		return true
	}
	return false
}
