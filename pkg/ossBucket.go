package pkg

import (
	"log"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func GetBucket(ossClient *oss.Client, bucketName string) *oss.Bucket {
	bucket, err := ossClient.Bucket(bucketName)
	if err != nil {
		log.Fatalf("获取 bucket 出错: %v\n", err)
	}

	return bucket
}
