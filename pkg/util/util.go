package util

import (
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





