package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func EncryptAes(origData, key []byte) ([]byte, error) {
	result, err := AesEncrypt(origData, key)
	if err != nil {
		return nil, err
	}
	_en_base64 := base64.StdEncoding.EncodeToString(result)
	return []byte(_en_base64), nil
}

func DecryptAes(crypted, key []byte) ([]byte, error) {
	_de_base64, err := base64.StdEncoding.DecodeString(string(crypted))
	if err != nil {
		return nil, err
	}
	origData, err := AesDecrypt(_de_base64, key)
	if err != nil {
		return nil, err
	}
	return origData, nil
}

func EncryptAesString(origData, key string) (string, error) {
	result, err := AesEncrypt([]byte(origData), []byte(key))
	if err != nil {
		return "", err
	}
	_en_base64 := base64.StdEncoding.EncodeToString(result)
	return _en_base64, nil
}

func DecryptAesString(crypted, key string) (string, error) {
	_de_base64, err := base64.StdEncoding.DecodeString(crypted)
	if err != nil {
		return "", err
	}
	origData, err := AesDecrypt(_de_base64, []byte(key))
	if err != nil {
		return "", err
	}
	return string(origData), nil
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}
