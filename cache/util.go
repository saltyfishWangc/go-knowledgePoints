package cache

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// 定义单位大小
const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)

func ParseSize(size string) (int64, string) {
	reg, _ := regexp.Compile("[0-9]+")
	uint := string(reg.ReplaceAll([]byte(size), []byte("")))
	num, _ := strconv.ParseInt(strings.Replace(size, uint, "", 1), 10, 64)
	uint = strings.ToUpper(uint)
	var byteNum int64 = 0
	switch uint {
	case "B":
		byteNum = num
	case "KB":
		byteNum = num * KB
	case "MB":
		byteNum = num * MB
	case "GB":
		byteNum = num * GB
	case "TB":
		byteNum = num * TB
	case "PB":
		byteNum = num * PB
	default:
		num = 0
		byteNum = 0
	}
	// 默认是100M
	if num == 0 {
		log.Println("ParseSize 仅支持 B KB MB GB TB PB")
		num = 100
		uint = "MB"
		byteNum = num * MB
	}
	sizeStr := strconv.FormatInt(num, 10) + uint
	return byteNum, sizeStr
}

// GetValSize 获取数据的大小
func GetValSize(val interface{}) int64 {
	bytes, _ := json.Marshal(val)
	size := int64(len(bytes))
	return size
}
