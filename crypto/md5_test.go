package crypto

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {

	_encodeMd5 := Md5("go语言中，判断两个字符串是否相等")
	fmt.Println(_encodeMd5)

	_bTrue := CompareMd5("go语言中，判断两个字符串是否相等", _encodeMd5)

	if !_bTrue {
		t.Errorf("CompareMd5 error")
	}
}
