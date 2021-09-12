package funcpro

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"log"
	"runtime/debug"
	"testing"
)

type TestObj struct {
	Name string
	Age  int
}

func TestGetKeyList(t *testing.T) {
	d := TestObj{"Jonham", 30}
	keys := getKeyList(d)
	assert.Equal(t, len(keys), 2)
	assert.Equal(t, keys[0], "Name")
	assert.Equal(t, keys[1], "Age")
}

func TestGetKeyListErrorType(t *testing.T) {
	defer func() {
		var err error
		if p := recover(); p != nil {
			str, ok := p.(string)
			if ok {
				err = errors.New(str)
			} else {
				err = errors.New("panic")
			}
			debug.PrintStack()
			log.Println(err.Error())
		}
	}()

	d := 12
	keys := getKeyList(d)
	log.Println(keys)
}
