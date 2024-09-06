# hexo-alioss
用于hexo的阿里云OSS对象存储上传工具

## 使用方法

1. 通过 release 下载二进制文件或下载源码编译
2. 参考 `example.json` 或下面的示例来配置你自己的 `config.json`
3. gopher 可以使用 `go install` 命令直接下载安装

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
- [ ] 添加断点续传
- [x] 添加 MD5 标示，不再上传已上传文件


