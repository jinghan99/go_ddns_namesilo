package handle

import (
	"log"
	"testing"
)

func Test_myIp(t *testing.T) {
	for i := 0; i < 5; i++ {
		log.Println(myIp())
	}
}

func TestDDnsByNameSilo(t *testing.T) {
	DDnsByNameSilo()
}

// 测试  ipv6 地址获取
func TestGetIPv6Address(t *testing.T) {
	address, _ := GetIPv6Address()
	log.Println(address)
}
