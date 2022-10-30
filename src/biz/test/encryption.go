package test

import (
	utils "github.com/bnc1010/containerManager/biz/utils"
	"fmt"
	"log"
)

func EncryptionTest() {
	// AES-CBC
	iv := []byte{0x19, 0x34, 0x57, 0x72, 0x90, 0xAB, 0xCD, 0xEF, 0x12, 0x64, 0x14, 0x78, 0x90, 0xAC, 0xAE, 0x45}
	key := []byte("1111111111111111")
	AesCbc := utils.CBCEncrypt([]byte("hello"), key, iv)
	fmt.Println(AesCbc)
	content := utils.CBCDecrypt(AesCbc, key, iv)
	fmt.Println(string(content))
	// BASE64
	encode := utils.Base64Encoding([]byte("Hello"))
	fmt.Println(encode)
	decode := utils.Base64Decoding(encode)
	fmt.Println(decode)
	// BASE58
	encode = utils.Base58Encoding("Hello")
	fmt.Println(encode)
	decode = utils.Base58Decoding(encode)
	fmt.Println(decode)
	// 哈希算法
	md4 := utils.HashMD4Encoding("hello")
	fmt.Println(md4)
	md5 := utils.HashMD5Encoding("hello")
	fmt.Println(md5)
	sha := utils.HashSHA256Encoding("hello")
	fmt.Println(sha)
	// DES
	encode, _ = utils.DesEncoding("hello", []byte("11111111")) // 8位数
	fmt.Println("Des:", encode)
	decode, _ = utils.DesDecoding(encode, []byte("11111111"))
	fmt.Println("Des:", decode)
	encode, _ = utils.TDesEncoding("hello", []byte("111111111111111111111111")) // 24位数
	fmt.Println(encode)
	decode, _ = utils.TDesDecoding(encode, []byte("111111111111111111111111"))
	fmt.Println(decode)
	// RSA
	err := utils.SaveRsaKey(2048)
	if err != nil {
		log.Println(err)
	}
	encoding, _ := utils.RsaEncoding("hello", "utilsKey.pem")
	log.Println(utils.Base64Encoding(encoding))
	decoding, _ := utils.RsaDecoding(encoding, "privateKey.pem")
	log.Println(string(decoding))
	// 数字签名
	encoding, _ = utils.RsaSign("privateKey.pem", "hello")
	log.Println(encoding)
	isTrue, _ := utils.RsaVerify(encoding, "hello", "utilsKey.pem")
	log.Println(isTrue)
}