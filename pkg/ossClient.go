package pkg

import (
	"log"
	"sync"

	"github.com/FarmerChillax/hexo-alios/vars"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	client *oss.Client
	once   sync.Once
)

func createOssClient() *oss.Client {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New(vars.AliOssConfig.Endpoint,
		vars.AliOssConfig.AccessKey, vars.AliOssConfig.AccessKeySecret)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	return client
}

func GetOssClient() *oss.Client {
	if client == nil {
		once.Do(func() {
			log.Printf("正在创建阿里云对象存储客户端")
			client = createOssClient()
		})
	}
	return client
}
