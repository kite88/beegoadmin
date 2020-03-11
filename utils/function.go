package utils

import (
	"beegoadmin/utils/storage"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"
)

/**
md5 加密
*/
func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

/**
通过反射获取数值
*/
func GetDataValue(fieldName string, data interface{}) interface{} {
	value := reflect.ValueOf(data)
	return value.FieldByName(fieldName).Interface()
}

/*
	生成[n.m]的随机数
*/
func Rand(startNum, endNum int) int {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rd.Intn(endNum-startNum+1) + startNum
}

// 获取本机网卡IP 内网IP
func GetLocalIP() (ipv4 string, err error) {
	var (
		addrs   []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet // IP地址
		isIpNet bool
	)
	// 获取所有网卡
	if addrs, err = net.InterfaceAddrs(); err != nil {
		return
	}
	// 取第一个非lo的网卡IP
	for _, addr = range addrs {
		// 这个网络地址是IP地址: ipv4, ipv6
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			// 跳过IPV6
			if ipNet.IP.To4() != nil {
				ipv4 = ipNet.IP.String() // 192.168.1.1
				return
			}
		}
	}
	return
}

//通过外网获取公网ip
func GetPublicNetIP() (ip string, err error) {
	resp, err0 := http.Get("http://ipv4.icanhazip.com/")
	if err0 != nil {
		err = err0
		return
	}
	defer resp.Body.Close()
	content, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		err = err1
		return
	}
	ip = strings.Trim(string(content), "\n")
	return
}

//通过dns服务器8.8.8.8:80获取使用的ip
func GetPulicIP() (ip string, err error) {
	conn, err := net.Dial("http", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	ip = localAddr[0:idx]
	return
}

//获取程序执行目录
func GetExecutePath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return dir, err
}

//获取计算机存储信息
func GetDisk() string {
	var bytes string
	if runtime.GOOS == "linux" {
		return ""
	}
	if runtime.GOOS == "windows" {
		info, _ := storage.GetStorageInfo()
		for _, val := range info {
			bytes += val.Name + "总共" + FormatByte(val.Total) + ",空闲" + FormatByte(val.Free) + "; "
		}
		return bytes
	}
	return ""
}

//字节单位转化
func FormatByte(bytes uint64) string {
	unit := [...]string{"B", "KB", "MB", "GB", "TB", "PB"}
	var i int
	for i = 0; bytes >= 1024 && i < len(unit)-1; i++ {
		bytes /= 1024
	}
	return fmt.Sprintf("%d%s", bytes, unit[i])
}

//时间格式化
func FormatDateTime(unixDateTime int64, format string) string {
	if 0 == unixDateTime {
		return ""
	}
	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	return time.Unix(unixDateTime, 0).Format(format)
}

//获取当天的时间戳 00:00:00 ~ 23:59:59 的两个边界时间戳
func GetSameDayUnix() (BeginUnix int64, EndUnix int64) {
	timeStr := time.Now().Format("2006-01-02")
	loc, _ := time.LoadLocation("Asia/Shanghai")
	BeginTimeParse, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 00:00:00", loc)
	EndTimeParse, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 23:59:59", loc)

	BeginUnix = BeginTimeParse.Unix()
	EndUnix = EndTimeParse.Unix()
	return
}

//struct转map
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

//struct转map(struct->json->map)
func StructToJsonToMap(obj interface{}) map[string]interface{} {
	//struct转json
	jsonBytes, _ := json.Marshal(obj)
	//json转map
	var mapResult map[string]interface{}
	json.Unmarshal([]byte(jsonBytes), &mapResult)
	return mapResult
}