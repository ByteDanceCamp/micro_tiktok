package utils

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"micro_tiktok/pkg/constants"
	"mime/multipart"
)

func UpLoadFile(file multipart.File, fileSize int64) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope:         constants.QiNiuBucket,
		PersistentOps: "vframe/jpg/offset/7",
	}
	mac := qbox.NewMac(constants.QiNiuAccessKey, constants.QiNiuSecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.Put(context.Background(), &ret, upToken, "aaa", file, fileSize, &putExtra)
	if err != nil {
		return "", err
	}

	url := constants.QiniuServer + ret.Key
	return url, nil
}
