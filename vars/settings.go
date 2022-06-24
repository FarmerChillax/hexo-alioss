package vars

import (
	"log"

	"github.com/spf13/viper"
)

type OssConfig struct {
	Endpoint        string
	AccessKey       string
	AccessKeySecret string
	Bucket          string
	Root            string
}

var AliOssConfig *OssConfig

func init() {
	viper.SetConfigName("config") //指定配置文件的文件名称(不需要制定配置文件的扩展名)
	//viper.AddConfigPath("/etc/appname/")   //设置配置文件的搜索目录
	//viper.AddConfigPath("$HOME/.appname")  // 设置配置文件的搜索目录
	viper.AddConfigPath(".")    // 设置配置文件和可执行二进制文件在用一个目录
	err := viper.ReadInConfig() // 根据以上配置读取加载配置文件
	if err != nil {
		log.Fatalf("读取配置文件出错: %v", err)
	}
	AliOssConfig = &OssConfig{
		Endpoint:        viper.GetString("endpoint"),
		AccessKey:       viper.GetString("ak"),
		AccessKeySecret: viper.GetString("aks"),
		Bucket:          viper.GetString("bucket"),
		Root:            viper.GetString("root"),
	}
	// fmt.Println("获取配置文件的string", viper.GetString(`app.name`))
	// fmt.Println("获取配置文件的string", viper.GetInt(`app.foo`))
	// fmt.Println("获取配置文件的string", viper.GetBool(`app.bar`))
	// fmt.Println("获取配置文件的map[string]string", viper.GetStringMapString(`app`))
}
