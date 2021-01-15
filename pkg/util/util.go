package util

import (
	"crypto/md5"
	"encoding/hex"
	"hash/crc32"
	"math/rand"
	"time"
)

func init() {
	// 以时间作为初始化种子
	rand.Seed(time.Now().UnixNano())
}

const (
	ascstr = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numstr = "1123456789"
)

// 获取随机字符TOKEN，输出字符数由参数size指定。
func GenNToken(size int) string {
	var bytes = make([]byte, size)
	rand.Read(bytes)
	length := byte(len(ascstr))
	for k, v := range bytes {
		bytes[k] = ascstr[v%length]
	}
	return string(bytes)
}

// 获取随机数字TOKEN，输出字符数由参数size指定。
func GenNumberToken(size int) string {
	var bytes = make([]byte, size)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = numstr[v%10]
	}
	return string(bytes)
}


func MD5(key string) string {
	ha := md5.New()
	ha.Reset()
	ha.Write([]byte(key))

	return hex.EncodeToString(ha.Sum(nil))
}


func GetHashCode(str string, count int) int {
	v := crc32.ChecksumIEEE([]byte(str))
	if v < 0 {
		v = -v
	}
	return int(v) % count
}


