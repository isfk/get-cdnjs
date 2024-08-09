package pkg

import (
	"bytes"
	"context"
	"fmt"

	"github.com/isfk/get-cdnjs/config"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

// 字节上传
func Upload(key string, data []byte) error {
	putPolicy := storage.PutPolicy{
		SaveKey: key,
		Scope:   fmt.Sprintf("%s:%s", config.Conf.Bucket, key),
	}

	mac := auth.New(config.Conf.AccessKey, config.Conf.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	zone, err := storage.GetZone(config.Conf.AccessKey, config.Conf.Bucket)
	if err != nil {
		return err
	}

	ret := &storage.PutRet{}
	err = storage.NewFormUploader(&storage.Config{Region: zone}).Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), int64(len(data)), &storage.PutExtra{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// 抓取网络文件
func Fetch(bucket, key, url string) error {
	_, err := storage.NewBucketManager(auth.New(config.Conf.AccessKey, config.Conf.SecretKey), &storage.Config{}).Fetch(url, bucket, key)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func List(bucket, prefix string) ([]string, error) {
	_, prefixes, _, _, err := storage.NewBucketManager(auth.New(config.Conf.AccessKey, config.Conf.SecretKey), &storage.Config{}).ListFiles(bucket, prefix, "/", "", 1000)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return prefixes, nil
}
