# hexo-alioss
用于hexo的阿里云OSS对象存储上传工具


## version

go >= 1.16

## 使用方法

1. 通过 release 下载二进制文件或下载源码编译
2. 参考 `example.json` 或下面的配置来配置你自己的 `config.json`

```json
{
    "endpoint": "oss-cn-shenzhen.aliyuncs.com",
    "ak": "YourAccessKey",
    "aks": "YourAccessKeySecret",
    "bucket": "YourBucketName",
    "root": "上传目录路径"
}
```

3. 运行二进制文件

## todos

- [ ] 添加命令行选项
- [ ] 添加多种配置文件类型
