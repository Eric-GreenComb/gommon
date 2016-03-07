package crypto

import (
	"testing"
)

func TestHash(t *testing.T) {

	_b, _ := GenerateHash("ministor")

	_is := CompareHash(_b, "ministor")

	if !_is {
		t.Errorf("CompareHash %b", _is)
	}
}
