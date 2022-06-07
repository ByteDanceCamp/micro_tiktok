package utils

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"strings"
	"time"
)

var writeSuffixList = map[string]struct{}{
	"mp4":  {},
	"avi":  {},
	"wma":  {},
	"mpeg": {},
	"mpg":  {},
	"mov":  {},
}

// ValidateVideoInfo 校验上传文件是否符合要求
// 若符合则返回编码后的文件名及后缀
func ValidateVideoInfo(header *multipart.FileHeader) (string, string, error) {
	fileNameOrg := header.Filename
	lastDotIndex := strings.LastIndex(fileNameOrg, ".")
	if lastDotIndex < 0 {
		return "", "", errors.New("miss suffix")
	}
	suffix := fileNameOrg[lastDotIndex+1:]
	suffix = strings.ToLower(suffix)
	if _, ok := writeSuffixList[suffix]; !ok {
		return "", "", errors.New("suffix is invalid")
	}
	m := md5.New()
	if _, err := io.WriteString(m, fileNameOrg[:lastDotIndex]+time.Now().String()); err != nil {
		return "", "", err
	}
	fileName := fmt.Sprintf("%x", m.Sum(nil))
	pre := time.Now().Format("06/01/02")
	return pre + "/" + fileName, "." + suffix, nil
}
