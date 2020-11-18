package aliyunoss

import (
	"crypto/md5"
	"strings"

	"encoding/hex"

	"hash/crc32"

	"github.com/satori/go.uuid"
)

func UUID() string {
	uid := uuid.NewV1()
	return uid.String()
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func Md5With16(str string) string {
	return Md5(str)[8:24]
}

func GetHashCode(str string, count int) int {
	v := crc32.ChecksumIEEE([]byte(str))
	if v < 0 {
		v = -v
	}
	return int(v) % count
}

func Dmd5(str string) string {
	return strings.ToUpper(Md5(str))
}
