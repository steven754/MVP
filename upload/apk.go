package upload

import (
	"MVP/pkg/file"
	"MVP/pkg/logging"
	"MVP/pkg/setting"
	"fmt"
	"os"
	"strings"
	"time"
)

//获取apk文件的保存路径，就是配置文件设置的   如：upload/apks/
func GetApkFilePath() string {
	return setting.AppSetting.ApkSavePath
}

//获取apk文件完整访问URL 如：http://127.0.0.1:8000/upload/apks/20190730/******.apk
func GetApkFullUrl(name string) string {
	return setting.AppSetting.ImagePrefixUrl + "/" + GetApkFilePath() + GetApkDateName() + name
}

//日期文件夹      如：20190730/
func GetApkDateName() string {
	t := time.Now()
	return fmt.Sprintf("%d%02d%02d/", t.Year(), t.Month(), t.Day())
}

//获取apk文件在项目中的目录  如：runtime/upload/apks/
func GetApkFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetApkFilePath()
}

//检查文件后缀，是否属于配置中允许的后缀名
func CheckApkExt(fileName string) bool {
	ext := file.GetExt(fileName)

	if strings.ToLower(ext) == strings.ToLower(setting.AppSetting.ApkAllowExt) {
		return true
	}

	return false
}

//检查apk文件
func CheckApk(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		logging.Warn("pkg/upload/apk.go文件CheckApk方法os.Getwd出错", err)
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src) //如果不存在则新建文件夹
	if err != nil {
		logging.Warn("pkg/upload/apk.go文件CheckApk方法file.IsNotExistMkDir出错", err)
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src) //检查文件权限
	if perm == true {
		logging.Warn("pkg/upload/apk.go文件CheckApk方法file.CheckPermission出错", err)
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
