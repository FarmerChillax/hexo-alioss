package main

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/FarmerChillax/hexo-alioss/pkg"
	"github.com/FarmerChillax/hexo-alioss/vars"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

func UploadFloder(bucket *oss.Bucket, rootPath string) {
	filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
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
			logrus.Errorf("bucket.GetObjectMeta err: %v", err)
			return err
		}
		if h.Get("Etag") == fmt.Sprintf("\"%s\"", fileMD5) {
			logrus.Infof("%s was uploaded, done.", objKey)
			return nil
		}

		err = bucket.PutObjectFromFile(objKey, path)
		if err != nil {
			color.Red("Upload %s err: %v\n", path, err)
			logrus.Errorf("Upload err: %v", err)
			return err
		}
		color.Green("Upload %s success.\n", path)
		return nil
	})
}

func main() {
	color.Yellow("OSS Go SDK Version: %v", oss.Version)
	// 初始化客户端
	ossClient := pkg.GetOssClient()
	bucket := pkg.GetBucket(ossClient, vars.AliOssConfig.Bucket)

	UploadFloder(bucket, vars.AliOssConfig.Path)
}
