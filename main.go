package main

import (
	"fmt"
	"github.com/robfig/cron"
	"go_ddns_namesilo/handle"
	"io"
	"log"
	"os"
)

func init() {
	// 追加日志 可添加 O_APPEND
	writer1, logErr := os.OpenFile("ddns.log", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if logErr != nil {
		fmt.Println("创建日志失败")
		os.Exit(1)
	}
	//os.Stdout代表标准输出流
	writer2 := os.Stdout
	// //io.MultiWriter实现多目的地输出 组合一下即可，
	multiWriter := io.MultiWriter(writer1, writer2)

	//	设置 logger 输出日志文件
	log.SetOutput(multiWriter)
	//	设置 logger 输出时间格式
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	//创建一个cron实例
	c := cron.New()
	//执行定时任务 每30分钟执行一次
	err := c.AddFunc("0 0/30 * * * ? ", func() {
		log.Println("=============DDnsByNameSilo 启动 每30分钟执行 =============")
		go handle.DDnsByNameSilo()
	})
	if err != nil {
		log.Println(err)
	}

	//启动定时任务
	c.Start()
	//退出程序时 关闭定时任务
	defer c.Stop()
	log.Println("=============ddns_namesolo 启动成功============= ")
	//一直保持程序运行
	select {}
}
