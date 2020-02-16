package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)
/**
	md5 加密
 */
func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str) )
	return fmt.Sprintf("%x", hash.Sum(nil))
}
/**
	通过反射获取数值
 */
func GetDataValue(fieldName string,data interface{}) interface{} {
	value := reflect.ValueOf(data)
	return value.FieldByName(fieldName).Interface()
}
/*
	生成[n.m]的随机数
 */
func Rand(startNum, endNum int) int {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rd.Intn(endNum - startNum + 1) + startNum
}
