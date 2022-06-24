package main

import (
	"io/fs"
	"log"
	"path/filepath"

	"github.com/FarmerChillax/hexo-alioss/pkg"
	"github.com/FarmerChillax/hexo-alioss/vars"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/fatih/color"
)

func UploadFloder(bucket *oss.Bucket, rootPath string) {
	filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		err = bucket.PutObjectFromFile(path, path)
		if err != nil {
			color.Red("Upload %s err: %v\n", path, err)
			log.Fatalf("Upload err: %v", err)
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

	UploadFloder(bucket, vars.AliOssConfig.Root)
}
