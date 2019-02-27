package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

// JSONTime json的格式化时间
type JSONTime time.Time

// Logger 输出到日志的变量
var Logger *log.Logger

func init() {
	file, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	Logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

// MarshalJSON 实现了json序列化JSONTime的方法
func (jsonTime JSONTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(jsonTime).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
