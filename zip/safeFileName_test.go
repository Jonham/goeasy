package zip

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestParseToSafeFileName(t *testing.T) {
	result := ParseToSafeFileName("hi/good//to go \\1233、上")
	logrus.Println(result)
}
