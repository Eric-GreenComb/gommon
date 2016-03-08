package uuid

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"strings"

	googleuuid "code.google.com/p/go-uuid/uuid"

	"github.com/banerwai/gommon/crypto"
)

// Google UUID
func UUID() string {
	return strings.Replace(googleuuid.NewRandom().String(), "-", "", -1)
}

// Guid
func NewGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return crypto.Md5(base64.URLEncoding.EncodeToString(b))
}

// 后面加个str生成之, 更有保障, 确保唯一
func NewGuidWith(str string) string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return crypto.Md5(base64.URLEncoding.EncodeToString([]byte(string(b) + str)))
}
