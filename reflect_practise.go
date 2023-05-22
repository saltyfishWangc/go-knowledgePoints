package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var logger *log.Logger = log.New(os.Stdout, "【wangc】", log.Llongfile|log.LstdFlags)

// 定义server结构体
type server struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}

// 定义mysql结构体
type mysql struct {
	Username string  `json:"username"`
	Passwd   string  `json:"passwd"`
	Database string  `json:"database"`
	Host     string  `json:"host"`
	Port     int     `json:"port"`
	Timeout  float64 `json:"timeout"`
}

// 定义配置结构体
type config struct {
	server `json:"server"`
	mysql  `json:"mysql"`
}

func mainReflectPractise() {
	//func main() {
	reflectPractise("./config/config.txt")
}

// reflectPractise 反射练习
/**
任务：解析如下配置文件
	1. 序列化：将结构体序列化为配置文件数据并保存到硬盘
	2. 反序列化：将配置文件内容反序列化到程序的结构体
配置文件内容如下：
#this is comment
;this a comment
;[]表示一个section
[server]
ip = 10.238.2.2
port = 8080

[mysql]
username = root
passwd = admin
database = test
host = 192.168.10.10
port = 8000
timeout = 1.2
*/
func reflectPractise(filePath string) {
	// 读取文件，一行一行的读
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	//// 定义server结构体
	//type server struct {
	//	Ip   string `json:"ip"`
	//	Port int    `json:"port"`
	//}
	//
	//// 定义mysql结构体
	//type mysql struct {
	//	Username string  `json:"username"`
	//	Passwd   string  `json:"passwd"`
	//	Database string  `json:"database"`
	//	Host     string  `json:"host"`
	//	Port     int     `json:"port"`
	//	Timeout  float64 `json:"timeout"`
	//}
	//
	//// 定义配置结构体
	//type config struct {
	//	server `json:"server"`
	//	mysql  `json:"mysql"`
	//}

	var serverConfig []string
	var mysqlConfig []string
	// 当前文本行是属于什么配置的，如果遇到[server]，直到再遇到别的配置，那currentParseTag的值一直是server，说明是server的配置
	var currentParseTag string

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			// 读完了
			break
		}
		if err != nil {
			panic(err)
		}
		if len(line) == 0 {
			// 空行不做处理，继续读
			logger.Println("空行不做处理，继续读")
			continue
		}

		// 去除收尾的空格
		content := strings.TrimSpace(string(line))

		// 过滤注释行
		if strings.HasPrefix(content, "#") || strings.HasPrefix(content, ";") {
			logger.Printf("【%s】 是注释行\n", content)
			continue
		}

		// 解析
		if strings.HasPrefix(content, "[server]") {
			// 初始化serverConfig切片，后面将相关的配置行都放在该切片中，最后统一映射到结构体中
			//serverConfig = make([]string, 10)
			currentParseTag = "server"
		} else if strings.HasPrefix(content, "[mysql]") {
			// 初始化mysqlConfig切片，后面将相关的配置行都放在该切片中，最后统一映射到结构体中
			//mysqlConfig = make([]string, 10)
			currentParseTag = "mysql"
		} else {
			switch currentParseTag {
			case "server":
				serverConfig = append(serverConfig, content)
			case "mysql":
				mysqlConfig = append(mysqlConfig, content)
			default:
				logger.Printf("【%s】 不属于任何配置\n", content)
			}
			continue
		}
	}

	configObj := &config{}

	// 开始序列化server、mysql
	if serverConfig != nil {
		serverObj := &server{}
		reflectStruct(serverConfig, serverObj)
		configObj.server = *serverObj
	}

	if mysqlConfig != nil {
		mysqlObj := &mysql{}
		reflectStruct(mysqlConfig, mysqlObj)
		configObj.mysql = *mysqlObj
	}

	logger.Printf("最后配置信息：%+v", configObj)

	//genConfigFile(configObj, "./config/outputConfig.txt")
	genConfigFile(*configObj, "./config/outputConfig.txt")
}

// reflectStruct 解析配置行反射到结构体对象
func reflectStruct(content []string, any any) {
	valueOf := reflect.ValueOf(any).Elem()
	typeOf := valueOf.Type()
	for _, content := range content {
		// 将文本行按照=切割
		split := strings.Split(content, "=")
		key := strings.TrimSpace(split[0])
		value := strings.TrimSpace(split[1])
		for i := 0; i < valueOf.NumField(); i++ {
			field := valueOf.Field(i)
			// 如果这个字段类型对象的tag中，json tag的key和配置对应，就设置字段值
			if key == typeOf.Field(i).Tag.Get("json") {
				// 根据字段类型赋值
				switch field.Type().Kind() {
				case reflect.Int:
					value, err := strconv.Atoi(value)
					if err != nil {
						panic(err)
					}
					field.SetInt(int64(value))
				case reflect.String:
					field.SetString(value)
				case reflect.Float64:
					value, err := strconv.ParseFloat(value, 64)
					if err != nil {
						panic(err)
					}
					field.SetFloat(value)
				}
			}
		}
	}
}

// genConfigFile 将配置序列化本地指定文件
func genConfigFile(any any, outputFilePath string) {
	// 如果传进来的any是个指针类型，那这里就必须用reflect.ValueOf(any).Elem()获取值信息，否则在获取typeOf.NumField()时会报错，Elem()是用来解引用的
	//valueOf := reflect.ValueOf(any).Elem()

	// 如果传进来的any不是个指针类型，那这里就用reflect.ValueOf(any)获取值信息
	valueOf := reflect.ValueOf(any)
	typeOf := valueOf.Type()
	// 配置文件内容
	var content string

	// 反射读取配置结构体的字段，判断每个字段的结构体类型
	for i := 0; i < typeOf.NumField(); i++ {
		tagOfField := typeOf.Field(i).Tag.Get("json")
		content = fmt.Sprintln(content, fmt.Sprintf("\n[%s]", tagOfField))
		// 下面开始解析字段对应的结构体
		valueOfField := valueOf.Field(i)
		typeOfField := valueOfField.Type()
		for j := 0; j < typeOfField.NumField(); j++ {
			var value string
			//switch typeOfField.Field(j).Type.Kind() {
			//case reflect.Float64:
			//	//value = strconv.FormatFloat(valueOfField.Field(j).Float(), 'f', 2, 64)
			//	value = fmt.Sprintf("%f", valueOfField.Field(j).Float())
			//case reflect.Int:
			//	value = strconv.Itoa(int(valueOfField.Field(j).Int()))
			//case reflect.String:
			//	value = valueOfField.Field(j).String()
			//default:
			//	logger.Println("不支持的类型:", value)
			//}
			//content = fmt.Sprintln(content, strings.TrimSpace(fmt.Sprintf("%s = %s", typeOfField.Field(j).Tag.Get("json"), value)))

			switch typeOfField.Field(j).Type.Kind() {
			case reflect.Float64:
				value = fmt.Sprintf("%0.2f", valueOfField.Field(j).Interface())
			case reflect.Int:
				value = fmt.Sprintf("%d", valueOfField.Field(j).Interface())
			case reflect.String:
				value = fmt.Sprintf("%s", valueOfField.Field(j).Interface())
			default:
				logger.Println("不支持的类型:", value)
			}
			content = fmt.Sprintln(content, strings.TrimSpace(fmt.Sprintf("%s = %s", typeOfField.Field(j).Tag.Get("json"), value)))
		}
	}

	file, err := os.Create(outputFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.Write([]byte(content))
	writer.Flush()
}
