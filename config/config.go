package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"syscall"
)

type Conf struct {
	NameSiloConf NameSiloConf `yaml:"nameSilo"`
}

// NameSiloConf 配置文件导入数据
type NameSiloConf struct {
	ApiKey   string `yaml:"apikey"`
	Domain   string `yaml:"domain"`
	DDnsHost string `yaml:"ddns_host"`
}

// MyConfig 初始化 全局配置参数
var MyConfig *NameSiloConf

func init() {
	//1 、初始化 yaml
	initYaml()
	//2、 初始化 系统配置参数环境变量
	initSysEnv()

	log.Printf("go_ddns_namesilo 配置：%v \n", MyConfig)

}

func initSysEnv() {
	// 环境变量中 有配置时
	ApiKey, found := syscall.Getenv("apikey")
	if found {
		MyConfig.ApiKey = ApiKey
	}
	domain, found := syscall.Getenv("domain")
	if found {
		MyConfig.Domain = domain
	}
	DDnsHost, found := syscall.Getenv("ddns_host")
	if found {
		MyConfig.DDnsHost = DDnsHost
	}
}

func initYaml() {
	//     ./是你当前的工程目录，并不是该go文件所对应的目录
	yamlFile, err := ioutil.ReadFile("./conf.yaml")
	if err != nil {
		log.Panicln(err)
	}
	var conf *Conf
	err = yaml.Unmarshal(yamlFile, &conf)

	if err != nil {
		log.Println(err)
	}
	//yaml 配置 赋值 至 配置结构体
	MyConfig = &conf.NameSiloConf
}
