package utils

import (
	"context"
	"encoding/base64"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"micro_tiktok/pkg/constants"
	"mime/multipart"
)

func UpLoadFile(file multipart.File, fileName, videoSuffix string, fileSize int64) (string, error) {
	coverKey := base64.StdEncoding.EncodeToString([]byte(constants.QiNiuBucket + ":cover/" + fileName + ".jpg"))
	putPolicy := storage.PutPolicy{
		Scope:         constants.QiNiuBucket,
		PersistentOps: "vframe/jpg/offset/1|saveas/" + coverKey,
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

	err := formUploader.Put(context.Background(), &ret, upToken, "video/"+fileName+videoSuffix, file, fileSize, &putExtra)
	if err != nil {
		return "", err
	}

	url := constants.QiNiuServer + ret.Key
	return url, nil
}
