package logging

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/moshushu/gin-blog/pkg/setting"
)

// 获取日志文件路径
func getLogFilePath() string {
	return setting.AppSetting.LogSavePath
}

// 获取完成的日志文件路径
func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission : %v", err)
	}
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}
	return handle
}

func mkDir() {
	// 返回与当前目录对应的根路径名
	dir, _ := os.Getwd()
	// 创建对应的目录以及所需的子目录
	// ModePerm =  0777
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
