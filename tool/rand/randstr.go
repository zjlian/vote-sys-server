package rand

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

const randomTarget string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// MD5 生成32位 参数字符串 的md5值
func MD5(text string) string {
	md5code := md5.New()
	md5code.Write([]byte(text))
	return hex.EncodeToString(md5code.Sum(nil))
}

// GetRandomString 生成 长度为 l 的伪随机字符串
func GetRandomString(l int) string {
	var bytes = []byte(randomTarget)
	var result = []byte{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < l; i++ {
		x := r.Intn(len(bytes))
		result = append(result, bytes[x])
	}

	return string(result)
}

// GetRS64 生成 长度为64的伪随机字符串
func GetRS64() string {
	return GetRandomString(64)
}

// GetRS32 生成 长度为32的伪随机字符串
func GetRS32() string {
	return GetRandomString(32)
}

// GetRS16 生成 长度为16的伪随机字符串
func GetRS16() string {
	return GetRandomString(16)
}

// GetRS8 生成 长度为8的伪随机字符串
func GetRS8() string {
	return GetRandomString(8)
}
