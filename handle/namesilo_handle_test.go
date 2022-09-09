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
