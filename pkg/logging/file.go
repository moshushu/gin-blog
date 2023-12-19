package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/moshushu/gin-blog/pkg/file"
	"github.com/moshushu/gin-blog/pkg/setting"
)

// 获取日志文件路径
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

// 获取完成的日志文件路径
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
}

func openLogFile(fileName, filePath string) (*os.File, error) {
	// 1、返回与当前目录对应的根路径名
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err :%v", err)
	}

	// 2、判断文件是否有权限
	src := dir + "/" + filePath
	perm := file.CheckPermission(src)
	if perm {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	// 3、判断文件路径是否存在
	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	// 4、是否能正常打开文件
	f, err := file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}
