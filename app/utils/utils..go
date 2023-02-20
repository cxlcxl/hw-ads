package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"github.com/shopspring/decimal"
	"math/rand"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// GetPages 获取分页参数
func GetPages(page, size int64) (offset int64) {
	offset = (page - 1) * size
	return
}

// CeilPages 计算总页数
func CeilPages(num, pageSize int64) int64 {
	if num < pageSize {
		return 1
	}
	var d int64 = 0
	if num%pageSize > 0 {
		d = 1
	}
	return num/pageSize + d
}

func MD5(params string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(params))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

// Base64 base64...
func Base64(params string) string {
	return base64.StdEncoding.EncodeToString([]byte(params))
}

// Shuffle 打乱数组原因顺序
func Shuffle(slice []interface{}) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

// GetFileExt 获取文件后缀
func GetFileExt(fp multipart.File) string {
	buffer := make([]byte, 32)
	if _, err := fp.Read(buffer); err != nil {
		// 获取失败
		return ""
	}
	return http.DetectContentType(buffer)
}

// GenerateSecret 生成密码加密串
func GenerateSecret(n int) string {
	if n == 0 {
		rand.Seed(time.Now().UnixNano())
		n = rand.Intn(15)
		if n < 3 {
			n = 8
		}
	}
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	bytes := []byte(str)
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = bytes[rand.Intn(len(bytes))]
	}
	return string(result)
}

// Password 登陆密码
func Password(pass, secret string) string {
	return strings.ToUpper(MD5(base64.StdEncoding.EncodeToString([]byte(secret + pass + secret))))
}

// InArray 判断是否在数组中
func InArray[T int64 | string](v T, vs []T) (ok bool) {
	if len(vs) > 50 {
		tmp := make(map[T]struct{})
		for _, t := range vs {
			tmp[t] = struct{}{}
		}
		_, ok = tmp[v]
	} else {
		for _, t := range vs {
			if t == v {
				ok = true
				break
			}
		}
	}
	return
}

// SomeInArray 多个元素一起判断是否在同一个数组中
func SomeInArray[T int64 | string](vs []T, v ...T) map[T]bool {
	tmp := make(map[T]struct{})
	for _, t := range vs {
		tmp[t] = struct{}{}
	}
	rs := make(map[T]bool)
	for _, t := range v {
		_, rs[t] = tmp[t]
	}
	return rs
}

// StringToFloat string to float64
func StringToFloat(d string) float64 {
	f, err := strconv.ParseFloat(d, 64)
	if err != nil {
		return 0
	}
	return f
}

// Round 保留小数位
func Round(f float64, n int) float64 {
	d, _ := decimal.NewFromFloat(f).Round(int32(n)).Float64()
	return d
}

// BufferConcat 字符串拼接
func BufferConcat(s []string, seq string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < len(s); i++ {
		if i > 0 && seq != "" {
			buf.WriteString(seq)
		}
		buf.WriteString(s[i])
	}
	return buf.String()
}

func WhereIn[T int64 | string](column string, v []T) (string, []interface{}) {
	conditions := make([]string, len(v))
	values := make([]interface{}, len(v))
	for i, t := range v {
		conditions[i] = "?"
		values[i] = t
	}
	return column + " in (" + strings.Join(conditions, ",") + ")", values
}
