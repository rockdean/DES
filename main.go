package main

import (
	"des"
	"encoding/base64"
	"fmt"
	"github.com/qiniu/iconv"
)

func main() {

	// DES 加解密
	testDes()
	// 3DES加解密
	test3Des()
}

func testDes() {
	cd, err := iconv.Open("gbk", "utf-8")
	if err != nil {
		fmt.Println("iconv.Open failed!")
		return
	}
	defer cd.Close()
	key := []byte("12345678")
	result, err := des.DesEncrypt([]byte(cd.ConvString("在那遥远的地方")), key)
	if err != nil {
		panic(err)
	}
	fmt.Println("DES :", base64.StdEncoding.EncodeToString(result))

	origData, err := des.DesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	cd, err = iconv.Open("utf-8", "gbk")
	if err != nil {
		fmt.Println("iconv.Open failed!")
		return
	}
	fmt.Println(cd.ConvString(string(origData)))
}

func test3Des() {
	key := []byte("sfe023f_sefiel#fi32lf3e!")
	result, err := des.TripleDesEncrypt([]byte("polaris@studygol"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := des.TripleDesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}
