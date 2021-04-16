package util

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/satori/go.uuid"
	"hash/crc32"
	"regexp"
	"strings"
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

func Dmd5(str string) string {
	return strings.ToUpper(Md5(str))
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

func StringDifference(a []string, b []string) []string {
	c := make([]string, len(a))
	for _, va := range a {
		hasVa := false
		for _, vb := range b {
			if va == vb {
				hasVa = true
				break
			}
		}
		if !hasVa {
			c = append(c, va)
		}
	}
	return c
}

// 比较两个字符串slice是否完全相等
func StringSliceEqual(sliceA, sliceB []string) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}

	if (sliceA == nil) != (sliceB == nil) {
		return false
	}

	for i, v := range sliceA {
		if v != sliceB[i] {
			return false
		}
	}

	return true
}

// 删除字符串slice中target值
func RemoveStringSliceTarget(strSlice *[]string, target string) {
	for index, v := range *strSlice {
		if v == target {
			*strSlice = append((*strSlice)[:index], (*strSlice)[index+1:]...)
			break
		}
	}
}

// 对字符串slice去重
func DistinctStringSlice(strSlice *[]string) {
	for index, v := range *strSlice {
		for j := index + 1; j < len(*strSlice); j++ {
			if v == (*strSlice)[j] {
				*strSlice = append((*strSlice)[:j], (*strSlice)[j+1:]...)
				j--
			}
		}
	}
}

//手机号验证
func ValitatorPhone(p string) bool {
	rex := `^(1(([35][0-9])|[8][0-9]|[9][0-9]|[6][0-9]|[7][01356789]|[4][579]))\d{8}$`
	reg := regexp.MustCompile(rex)
	return reg.MatchString(p)
}

//异常处理
func Try(block func(), catch func(e interface{}), finally func()) {
	defer func() {
		e := recover() // 当block中panic时，recover将返回panic的参数
		if e != nil {  // 当block中没有panic时，recover将返回nil
			catch(e)
		}
		if finally != nil {
			finally()
		}
	}()
	block() // 当block中panic时，recover将返回panic的参数
}
