package main

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/FarmerChillax/hexo-alioss/pkg"
	"github.com/FarmerChillax/hexo-alioss/vars"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

var Version = "v0.2.4"

type UploadFileItem struct {
	ObjectKey string
	FilePath  string
}

func UploadFloder(bucket *oss.Bucket, rootPath string, filePathChan chan<- *UploadFileItem) {
	filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		objKey, err := filepath.Rel(rootPath, path)
		if err != nil {
			color.Yellow("filepath.Rel err: %v", err)
			logrus.Errorf("filepath.Rel err: %v", err)
			return err
		}

		fileMD5 := pkg.GetMD5(path)
		h, err := bucket.GetObjectMeta(objKey)
		if err != nil {
			ossErr, ok := err.(oss.ServiceError)
			if ok && ossErr.StatusCode == 404 {
				logrus.Infof("upload a new file: %s", objKey)
			} else {
				logrus.Errorf("bucket.GetObjectMeta err: %v, objKey: %s", err, objKey)
				return err
			}
		}

		if h.Get("Etag") == fmt.Sprintf("\"%s\"", fileMD5) {
			logrus.Infof("%s was uploaded, done.", objKey)
			return nil
		}

		filePathChan <- &UploadFileItem{
			ObjectKey: objKey,
			FilePath:  path,
		}
		return nil
	})
}

func OSSUploader(ctx context.Context, bucket *oss.Bucket, filePathChan <-chan *UploadFileItem) {
	for {
		select {
		case <-ctx.Done():
			return
		case item, ok := <-filePathChan:
			if !ok {
				logrus.Infof("filePathChan is closed.")
				return
			}
			objKey := item.ObjectKey
			path := item.FilePath
			err := bucket.PutObjectFromFile(objKey, path)
			if err != nil {
				color.Red("Upload %s err: %v\n", path, err)
				logrus.Errorf("Upload err: %v", err)
			} else {
				color.Green("Upload %s success.\n", path)
			}
		}
	}
}

func main() {
	color.Green("hexo-alioss cli Version: %v", Version)
	color.Yellow("OSS Go SDK Version: %v", oss.Version)
	// 初始化客户端
	ossClient := pkg.GetOssClient()
	bucket := pkg.GetBucket(ossClient, vars.AliOssConfig.Bucket)

	filePathChan := make(chan *UploadFileItem, 8)

	for i := 0; i < cap(filePathChan); i++ {
		go OSSUploader(context.Background(), bucket, filePathChan)
	}

	UploadFloder(bucket, vars.AliOssConfig.Path, filePathChan)
}
