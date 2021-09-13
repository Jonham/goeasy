package yaml

import (
	"log"
	"testing"
)

func TestSplitRow(t *testing.T) {
	raw := `AliyunOSS:      
  AccessKeyId: KEYNAME  

  AccessKeySecret: Secret`

	book := ParseToBook(raw)
	for key, value := range book {
		log.Println(key, value)
	}
}
