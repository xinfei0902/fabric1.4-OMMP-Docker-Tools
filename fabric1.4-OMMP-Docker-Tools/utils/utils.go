package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

//ParseInt 转int
func ParseInt(i interface{}) (val int) {
	if value, ok := i.(float64); ok {
		return int(value)
	}
	val64, _ := strconv.ParseInt(fmt.Sprintf("%v", i), 10, 64)
	val = int(val64)
	return val
}

//ParseInt64 转int64
func ParseInt64(i interface{}) (val int64) {
	if value, ok := i.(float64); ok {
		return int64(value)
	}
	val, _ = strconv.ParseInt(fmt.Sprintf("%v", i), 10, 64)
	return val
}

//ParseFloat64 转float64
func ParseFloat64(i interface{}) (val float64) {
	val64, _ := strconv.ParseFloat(fmt.Sprintf("%v", i), 64)
	return val64
}

//ParseString 转string
func ParseString(i interface{}) string {
	if i == nil {
		return ""
	}
	val := fmt.Sprintf("%v", i)
	return val
}

//StructToMap Struct转成Map
func StructToMap(obj interface{}) map[string]interface{} {
	objMap := make(map[string]interface{})
	var v reflect.Value
	if reflect.TypeOf(obj).Kind() == reflect.Ptr {
		v = reflect.ValueOf(obj).Elem()
	} else {
		v = reflect.ValueOf(obj)
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		jsonname := t.Field(i).Tag.Get("json")
		if jsonname == "" || jsonname == "-" || jsonname == "_" {
			continue
		}
		objMap[jsonname] = v.Field(i).Interface()
	}
	return objMap
}

//SHA1 SHA1
func SHA1(str string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(str)))
}

//SHA256 SHA256
func SHA256(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}

//MD5 MD5
func MD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

//BASE64EncodeString base64 encode string
func BASE64EncodeString(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

//BASE64DecodeString base64 decode string
func BASE64DecodeString(str string) string {
	result, _ := base64.StdEncoding.DecodeString(str)
	return string(result)
}

//BASE64UrlEncodeString base64 encode string
func BASE64UrlEncodeString(str string) string {
	str = HmacMd5Encode(str)
	return base64.URLEncoding.EncodeToString([]byte(str))
}

//HmacEncode
func HmacMd5Encode(str string) string {
	var key string
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(str))
	return hex.EncodeToString(hmac.Sum([]byte("")))
}

//Round float精度
func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

//RandomString 随机字符
func RandomString(chars string, length int) string {
	rand.Seed(time.Now().Unix())
	str := []byte("")
	for i := 0; i < length; i++ {
		str = append(str, chars[rand.Intn(len(chars))])
	}
	return string(str)
}

func GetPreMonthString() string {
	return ParseString(int(time.Now().AddDate(0, -1, 0).Month()))
}
func GetMonthString() string {
	return ParseString(int(time.Now().Month()))
}

func GetTodayZeroTime() int64 {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local).Unix()
}

func GetYesterdayZeroTime() int64 {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1).Unix()
}

func GetMonthFirstDayZeroTime() int64 {
	y, m, _ := time.Now().Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, time.Local).Unix()
}
func GetMonthLastDay() time.Time {
	y, m, _ := time.Now().Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, time.Local).AddDate(0, 1, -1)
}
func GetPreMonthFirstDayZeroTime() int64 {
	//获取当前年月
	y, m, _ := time.Now().Date()
	//获取这个月第一天
	firstTime := time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
	//获取上个月第一天时间
	return firstTime.AddDate(0, -1, 0).Unix()
}

func GetTreeMonthAgoFirstDayZeroTime() int64 {
	//获取当前年月
	y, m, _ := time.Now().Date()
	//获取这个月第一天
	firstTime := time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
	//获取三个月前第一天时间
	return firstTime.AddDate(0, -3, 0).Unix()
}

func RemoveRepeatInt64(data []int64) []int64 {
	result := []int64{}
	tempMap := map[int64]bool{}
	for _, v := range data {
		if tempMap[v] == false {
			result = append(result, v)
			tempMap[v] = true
		}
	}
	return result
}
func RemoveRepeatString(data []string) []string {
	result := []string{}
	tempMap := map[string]bool{}
	for _, v := range data {
		if tempMap[v] == false {
			result = append(result, v)
			tempMap[v] = true
		}
	}
	return result
}

func HMACSHA256(plain string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(plain))
	return hex.EncodeToString(h.Sum(nil))
	// return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func MD5String(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
