## crypto

###使用方法：

- DES,3DES,Aes

```

package main

import (
	"encoding/base64"
	"fmt"
	"github.com/banerwai/gommon/crypto"
)

func main() {
	// DES 加解密
	testDes()
	// 3DES加解密
	test3Des()

	testAes()
}

func testDes() {
	key := []byte("sfe023f_")
	result, err := crypto.DesEncrypt([]byte("polaris@studygolang"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := crypto.DesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

func test3Des() {
	key := []byte("sfe023f_sefiel#fi32lf3e!")
	result, err := crypto.TripleDesEncrypt([]byte("polaris@studygolang"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := crypto.TripleDesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

func testAes() {
	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	key := []byte("sfe023f_9fd&fwflsfe023f_9fd&fwfl")
	// key := []byte("sfe023f_9fd&fwfl")
	result, err := crypto.AesEncrypt([]byte("polaris@studygolang"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := crypto.AesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

```